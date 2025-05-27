/*
	https://leetcode.com/problems/solving-questions-with-brainpower/description/

	You are given a 0-indexed 2D integer array questions where questions[i] =
		[pointsi, brainpoweri].

	The array describes the questions of an exam, where you have to process the
		questions in order
	(i.e., starting from question 0) and make a decision whether to solve or skip
		each question.
	Solving question i will earn you pointsi points but you will be unable to solve
		each of the
	next brainpoweri questions. If you skip question i, you get to make the
		decision on the next question.

		For example, given questions = [[3, 2], [4, 3], [4, 4], [2, 5]]:
			If question 0 is solved, you will earn 3 points but you will be unable to
				solve questions 1 and 2.
			If instead, question 0 is skipped and question 1 is solved, you will earn 4
				points but you will be

	unable to solve questions 2 and 3.

	Return the maximum points you can earn for the exam.
*/

package solutions

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n)
	dp[len(dp)-1] = questions[n-1][0]

	for i := n - 2; i >= 0; i-- {
		var alt int
		nextIdx := questions[i][1] + i + 1
		if nextIdx <= n-1 {
			alt = dp[nextIdx]
		}
		dp[i] = max(dp[i+1], questions[i][0]+alt)
	}

	return int64(dp[0])
}
