package work

import (
	"fmt"

	"github.com/atotto/clipboard"
	"clitut/network"
	//"clitut/file"
	"encoding/json"
	"clitut/schema"
	"time"
)

var localClip, localPreviousClip, ServerClip string

func Execute() {
	text, err := clipboard.ReadAll()
	if err == nil {
		//fmt.Println("Clipboard Data Local:")
		//fmt.Println(text)
		localPreviousClip = text
		fetch()

	} else {
		fmt.Println("Inside error copy")
		fmt.Println(err)
	}
}

/*func start()  {

	for {
		delaySecond(5)
	}

}*/

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}

func fetch() {

	localPreviousClip = localClip
	text, err := clipboard.ReadAll()
	if err == nil {
		localClip = text

		//	Now perform an api call
		resp, fetcherr := network.Fetch()

		if resp != nil {
			//m := new(schema.ServerArray)
			//fmt.Println(resp)

			fmt.Println("Going to Fetch Data")
			var temps []schema.Serverclip
			if err := json.NewDecoder(resp.Body).Decode(&temps); err != nil {

				fmt.Println(err)
				panic(err)
			} else {
				ServerClip = temps[0].Text

				task()

			}
		} else {
			fmt.Println("inside fetch error callback")
			fmt.Println(fetcherr)
			panic(fetcherr)
		}

	}
}

func task() {

	if localClip != localPreviousClip && localClip != ServerClip {

		/*Upload the data to server*/
		text, err := clipboard.ReadAll()
		if err == nil {

			network.Upload(text)
		}

	} else if localClip == localPreviousClip {

		localPreviousClip = localClip
		localClip = ServerClip
		clipboard.WriteAll(ServerClip)

	}

	delaySecond(5)
	fetch()

}

/*

var task=function(){


var newclip=clipboard.get('text');

console.log('running task function before if conditon');

if(localClip!=localPreviousClip && localClip!=serverClip){

console.log('running task function inside if conditon');
var data=Restangular.oneUrl('data','http://128.199.89.41:9001/data');

var params={user:VerifyLoginService.getUserID(),text:newclip};

data.save(params)
.then(function(res){
// alert('sent data to server');
console.log('posted');
console.log(res);

localClip=newclip;
localPreviousClip=localClip;
// $timeout(fetch,1000);
});

}

else if(localClip==localPreviousClip) {

localPreviousClip=localClip;
localClip=serverClip;
clipboard.set(serverClip, 'text');

// $timeout(fetch,1000);

}

$timeout(fetch,1000);

}
*/

