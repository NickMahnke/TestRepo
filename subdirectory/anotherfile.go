stuff

resp, gAPIErr := grequests.Get(fmt.Sprintf("https://att.com/"), &grequests.RequestOptions{InsecureSkipVerify: true})

resp, gAPIErr := grequests.Get(fmt.Sprintf("https://pge.com/"), &grequests.RequestOptions{InsecureSkipVerify: true})

resp, gAPIErr := grequests.Get(fmt.Sprintf("https://httpbin.org/anything?ssn=12345678"), &grequests.RequestOptions{InsecureSkipVerify: true})
