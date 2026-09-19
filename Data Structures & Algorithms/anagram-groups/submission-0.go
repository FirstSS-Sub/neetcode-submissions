func groupAnagrams(strs []string) [][]string {
    output := make([][]string, 0, len(strs))
    judgeMap := make(map[[26]int][]string)
    for _, s := range strs {
        var judge [26]int
        for _, r := range s {
            judge[int(r - 'a')] += 1
        }
        judgeMap[judge] = append(judgeMap[judge], s)
    }
    for _, v := range judgeMap {
        output = append(output, v)
    }
    return output
}
