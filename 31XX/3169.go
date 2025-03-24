package leetcode_ans

func countDays(days int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][0] < meetings[j][0]
    })

    mergedMeetings := [][]int{}
    for _, meeting := range meetings {
        if len(mergedMeetings) == 0 || meeting[0] > mergedMeetings[len(mergedMeetings)-1][1] {
            mergedMeetings = append(mergedMeetings, meeting)
        } else {
            mergedMeetings[len(mergedMeetings)-1][1] = max(mergedMeetings[len(mergedMeetings)-1][1], meeting[1])
        }
    }

    meetingDaysCount := 0
    for _, m := range mergedMeetings {
        meetingDaysCount += (m[1] - m[0] + 1)
    }

    return days - meetingDaysCount
}