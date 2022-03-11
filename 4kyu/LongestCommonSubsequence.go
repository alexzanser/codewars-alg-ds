package kata

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
  x, y := []byte(s1), []byte(s2)
  if len(x) == 0 || len(y) == 0{
    return ""
  }
  if x[len(x) - 1] == y[len(y) - 1] {
    return string(append([]byte(LCS(string(x[0:len(x) - 1]), string(y[0:len(y) - 1]))), x[len(x) - 1]))
  } else {
    left := LCS(string(x[0:len(x) - 1]), string(y))
    right:= LCS(string(x), string(y[0:len(y) - 1]))

    if len(left) > len(right) {
      return string(left)
    }else {
      return string(right)
    }
  }
}