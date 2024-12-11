using System.Numerics;

string input = "77 515 6779622 6 91370 959685 0 9861";
IEnumerable<BigInteger> BigIntegers = input.Split(" ").Select(s => BigInteger.TryParse(s, out BigInteger value) ? value : -1);

Console.WriteLine(part1(BigIntegers));
static BigInteger part1(IEnumerable<BigInteger> input)
{
    return NumberOfStones(input.ToList(), 25);
}

static BigInteger NumberOfDigits(BigInteger n)
{
    return n.ToString().Length;
}

static BigInteger NumberOfStones(List<BigInteger> stones, BigInteger iterations)
{
    for (BigInteger i = 0; i < iterations; i++)
    {
        BigInteger initialCount = stones.Count;
        for (BigInteger j = 0; j < initialCount; j++)
        {
            if (stones[(int)j] == 0)
            {
                stones[(int)j] = 1;
                continue;
            }
            if (NumberOfDigits(stones[(int)j]) % 2 == 0)
            {
                (BigInteger a, BigInteger b) = GetLeftRightDigits(stones[(int)j]);
                stones[(int)j] = a;
                stones.Add(b);
                continue;
            }
            stones[(int)j] *= 2024;
        }
    }
    return stones.Count;
}

static (BigInteger, BigInteger) GetLeftRightDigits(BigInteger n)
{
    string numString = n.ToString();
    int halfway = numString.Length / 2;
    string two = numString.Substring(halfway, halfway).TrimEnd('0').Length == 0 ? "0" : numString.Substring(halfway, halfway);
    return (BigInteger.TryParse(numString.Substring(0, halfway), out BigInteger v1) ? v1 : throw new Exception(), BigInteger.TryParse(two, out BigInteger v2) ? v2 : throw new Exception());
}

