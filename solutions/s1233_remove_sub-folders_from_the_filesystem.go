/*
	https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

	Given a list of folders folder, return the folders after removing all
		sub-folders in those folders. You may return the answer in any order.

	If a folder[i] is located within another folder[j], it is called a sub-folder
		of it. A sub-folder of folder[j] must start with folder[j], followed by a "/".
		For example, "/a/b" is a sub-folder of "/a", but "/b" is not a sub-folder of
		"/a/b/c".

	The format of a path is one or more concatenated strings of the form: '/'
		followed by one or more lowercase English letters.

		For example, "/leetcode" and "/leetcode/problems" are valid paths while an
			empty string and "/" are not.
*/

package solutions

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	res := []string{folder[0]}

	for i := 1; i < len(folder); i++ {
		last := res[len(res)-1] + "/"
		if !strings.HasPrefix(folder[i], last) {
			res = append(res, folder[i])
		}
	}

	return res
}
