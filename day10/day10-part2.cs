#:package Microsoft.Z3@4.12.2

using System.Text.RegularExpressions;
using Microsoft.Z3;

if (args.Length < 1) {
    Console.WriteLine("Requires input file");
}

var machines = ParseInputFile(args[0]);
var totalPresses = 0;

foreach (var machine in machines)
{
    using var ctx = new Context();
    
    // Create a variable for each button
    var vars = new List<IntExpr>();
    for (var i = 0; i < machine.Buttons.Count; i++)
    {
        vars.Add(ctx.MkIntConst($"x{i}"));
    }

    var solver = ctx.MkSolver();

    // All variables must be positive or 0
    foreach (var v in vars)
    {
        solver.Assert(ctx.MkGe(v, ctx.MkInt(0)));
    }

    // Loop through each joltage digit place
    for (var i = 0; i < machine.RequiredJoltage.Count; i++)
    {
        // Get all the button digits in this digit place
        var btnDigits = machine.Buttons.Select(btn => ctx.MkInt(btn[i]));

        // Sum of the all (button digit * var) should equal the joltage digit
        var digitSum = ctx.MkAdd(
            btnDigits.Select((d, i) => ctx.MkMul(vars[i], d))
        );

        solver.Assert(ctx.MkEq(digitSum, ctx.MkInt(machine.RequiredJoltage[i])));
    }

    int? minPresses = null;

    Status status;

    do {
        status = solver.Check();

        if (status == Status.SATISFIABLE)
        {
            Model model = solver.Model;

            // Determine how many times a button was pressed
            var presses = 0;

            foreach (var v in vars)
            {
                model.Evaluate(v);
                presses += int.Parse(model.Evaluate(v).ToString());
            }

            if (!minPresses.HasValue || presses < minPresses)
            {
                minPresses = presses;
            }

            // Add solution to exclusion list to force another solution
            solver.Assert(ctx.MkOr(
                vars.Select(v => ctx.MkNot(ctx.MkEq(v, model.Evaluate(v))))
            ));
        } 
    } while (status == Status.SATISFIABLE);

    totalPresses += minPresses ?? 0;
}

Console.WriteLine($"Part 2: {totalPresses}");

IEnumerable<Machine> ParseInputFile(string filename)
{
    var machines = new List<Machine>();
    var lines = File.ReadAllLines(filename);

    var joltageRx = new Regex(@"\{([\d,]+)\}");
    var buttonRx = new Regex(@"\(([\d+,]+)\)");

    foreach (var line in lines)
    {
        var joltageMatch = joltageRx.Match(line);
        var joltage = joltageMatch.Groups[1].Value.Split(",").Select(j => int.Parse(j)).ToList();

        var buttons = new List<List<int>>();
        var matches = buttonRx.Matches(line);
        foreach (Match match in matches)
        {
            var buttonPositions = match.Groups[1].Value.Split(",").Select(i => int.Parse(i));
            var button = new List<int>();
            
            for (var i = 0; i < joltage.Count ; i++)
            {
                if (buttonPositions.Contains(i))
                {
                    button.Add(1);
                } else
                {
                    button.Add(0);
                }
            }

            buttons.Add(button);
        }

        var machine = new Machine
        {
            Buttons = buttons,
            RequiredJoltage = joltage
        };

        machines.Add(machine);
    }

    return machines;
}

class Machine
{
    public IReadOnlyList<IReadOnlyList<int>> Buttons { get; set; } = [];

    public IReadOnlyList<int> RequiredJoltage { get; set; } = [];
}