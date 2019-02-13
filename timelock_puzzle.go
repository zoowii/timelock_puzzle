package timelock_puzzle

import (
	"math/big"
	"math/rand"
	crand "crypto/rand"
	"time"
	"github.com/pkg/errors"
)

var randomReader = rand.New(rand.NewSource(time.Now().UnixNano()))

var DefaultBits = 2048

func GeneratePrime(bits int) (*big.Int, error) {
	value, err := crand.Prime(randomReader, bits)
	return value, err
}

func GenerateRandomBigInt(bits int) (*big.Int, error) {
	return crand.Prime(randomReader, bits)
}

type PuzzleInfo struct {
	N *big.Int
	P *big.Int
	Q *big.Int
	T *big.Int
	X *big.Int
}

func GenerateTimeLockPuzzle(bits int, t *big.Int, x *big.Int) (result *PuzzleInfo, err error) {
	result = new(PuzzleInfo)
	p, err := GeneratePrime(bits)
	if err != nil {
		return
	}
	pStr := p.Text(10)
	result.P = p
	var q *big.Int
	for ;; {
		q, err = GeneratePrime(bits)
		if err != nil {
			return
		}
		if pStr != q.Text(10) {
			break
		}
	}
	result.Q = q
	result.N = big.NewInt(0)
	result.N = result.N.Mul(p, q)
	result.T = t
	result.X = x
	return
}

// 计算a^b
func bigIntPow(a *big.Int, b *big.Int) (*big.Int, error) {
	result := big.NewInt(1)
	// TODO: 改成O(log N)时间的算法
	one := big.NewInt(1)
	for i:=big.NewInt(0);i.Cmp(b)<0;i=i.Add(i, one) {
		result = result.Mul(result, a)
	}
	return result, nil
}

// 计算 2^t
func calculate2RaisedTo(t *big.Int) (result *big.Int, err error) {
	two := big.NewInt(2)
	return bigIntPow(two, t)
}

// 知道p和q的情况下计算y
// Phi(N) = Phi(p)*Phi(q)=(p-1)*(q-1)
// 计算 e = 2 ^ T (mod Phi(N))
// 然后计算 y = x ^ e ( mod N )
// 因为 e = 2 ^ T (mod Phi(N))，也就是 2^T = k*Phi(N) + e
// 根据欧拉定理, x ^ Phi(N) = 1 (mod N)，则 x^(2^T) = x^(k*Phi(N)+e) = ((x^(Phi(N))^k) * x^e = 1^k * x^e = x^1 = y (mod N)，满足time lock puzzle条件，只需要O(log(T)+log(Phi(N))次顺序乘法运算即可快速计算得到结果y
func SolveByPrivate(puzzle *PuzzleInfo) (y *big.Int, err error) {
	p := puzzle.P
	q := puzzle.Q
	t := puzzle.T
	x := puzzle.X
	n := puzzle.N
	if p == nil || q == nil {
		err = errors.New("require p and q")
		return
	}
	one := big.NewInt(1)
	phiN := big.NewInt(0)
	pm1 := big.NewInt(0)
	pm1 = pm1.Sub(p, one)
	qm1 := big.NewInt(0)
	qm1 = qm1.Sub(q, one)
	phiN = phiN.Mul(pm1, qm1) // Phi(N) = (p-1) * (q-1)
	raise2ToT, err := calculate2RaisedTo(t)
	if err != nil {
		return
	}
	e := big.NewInt(0)
	e = e.Mod(raise2ToT, phiN) // e = 2^T (mod Phi(N))
	xRaiseToE, err := bigIntPow(x, e)
	if err != nil {
		return
	}
	y = big.NewInt(0)
	y = y.Mod(xRaiseToE, n)
	return
}

// 不知道p和q的情况下计算y
// 需要计算 y = x^(2^T) (mod N)，需要O(T)次顺序乘法运算才能得到结果y，并且无法用多CPU来加快计算
func SolveByPublic(puzzle *PuzzleInfo) (y *big.Int, err error) {
	raise2ToT, err := calculate2RaisedTo(puzzle.T)
	if err != nil {
		return
	}
	xRaiseTo2RaiseToT, err := bigIntPow(puzzle.X, raise2ToT)
	if err != nil {
		return
	}
	y = big.NewInt(0)
	y = y.Mod(xRaiseTo2RaiseToT, puzzle.N)
	return y, nil
}

