// gcc -o test test.c httpclient.a -framework Foundation -framework AppKit -framework Security
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "httpclient.h"


int main()
{
	struct httpclient_post_with_sni_return result;

	result = httpclient_post_with_sni("https://www.baidu.com/test", "14.215.177.38:443", "www.baidu.com", "test123");
	if (result.r0 == -1) {
		printf("post failed, err: %s\n", result.r1);
		return 0;
	}

	printf("post success, code: %lld, resp_body: %s\n", result.r0, result.r1);
	return 0;
}
