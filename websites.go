package main

/**
* function to check if the website is active
 */
func getWebsites() []Website {
	websites := []Website{
		{
			name:           "google",
			url:            "https://google.com",
			expectedStatus: 200,
			interval:       5,
			reportToEmail:  "bemijonathan@gmail.com",
		},
		{
			name:           "bemijonathan.blog",
			url:            "https://jonathanatieneblog.netlify.app",
			expectedStatus: 200,
			interval:       5,
			reportToEmail:  "atienejonathan@gmail.com",
		},
	}

	return websites
}
