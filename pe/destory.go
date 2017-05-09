package pe

/*
DestoryProcess can destory process no matter the status running/ready/blocked
*/
// func (p *Process) DestoryProcess() (q Queue) {
// 	if p.Status == "running" {
// 		log.Printf("The running process %d has been destory", p.PID)
// 		return nil
// 	}
// 	q = remove(*p, q)
// 	log.Printf("The %s process %d has been destory", p.Status, p.PID)
// 	return
// }

// func remove(p Process, s Queue) (q Queue) {
// 	// return append(q[:i], q[i+1:]...)
// 	for k, v := range s {
// 		if v == p {
// 			s = append(s[:k], s[k+1:]...)
// 		}
// 	}
// 	q = s
// 	return
// }
