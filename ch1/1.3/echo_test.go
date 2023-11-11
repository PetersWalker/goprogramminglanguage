// Exercise 1.3: Experiment to measure the dif ference in running time bet ween our pot ent ial ly
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates par t of the
// time package, and Sec tion 11.4 shows how to write benchmark tests for systematic performance evaluation.)
package echo_test

import (
  "testing"
  "os/exec"
)

func BenchmarkEcho1(b *testing.B) {
  cmd := exec.Command("go run ./ch1/echo/echo1.go", "hello world hello world hello world hello world hello world")

  for i := 0; i < b.N; i++ {
    cmd.Output()

  }
}

func BenchmarkEcho2(b *testing.B) {
  cmd := exec.Command("go run ./ch1/echo/echo2.go", "hello world hello world hello world hello world hello world")

  for i := 0; i < b.N; i++ {
    cmd.Output()

  }
}

func BenchmarkEcho3(b *testing.B) {
  cmd := exec.Command("go run ./ch1/echo/echo3.go", "hello world hello world hello world hello world hello world")

  for i := 0; i < b.N; i++ {
    cmd.Output()

  }
}