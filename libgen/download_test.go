// Copyright © 2020 Ryan Ciehanski <ryan@ciehanski.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libgen

import (
	"strings"
	"testing"
)

func TestDownloadBook(t *testing.T) {
	book, err := GetDetails(
		[]string{"2F2DBA2A621B693BB95601C16ED680F8"},
		GetWorkingMirror(SearchMirrors),
		false,
		false,
		"",
	)
	if err != nil {
		t.Error(err)
	}

	if err := GetDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}
	if err := DownloadBook(book[0], ""); err != nil {
		t.Error(err)
	}
}

func TestGetDownloadURL(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := GetDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}

	if book[0].URL == "" {
		t.Error("URL empty")
	}
	if !strings.Contains(book[0].URL, "http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=") {
		t.Error("incorrect URL returned")
	}
}

func TestGetBokDownloadURL(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := getBokDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}

	if book[0].URL == "" {
		t.Error("no valid url found")
	}
}

func TestGetBooksdlDownloadURL(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := getBooksdlDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}
	if !strings.Contains(book[0].URL, "http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=") {
		t.Errorf("got: %s, expected: http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=", book[0].URL)
	}
}

func TestGetHref(t *testing.T) {
	results := getHref(booksdlReg, `<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'>
<html xmlns='http://www.w3.org/1999/xhtml'>
<head>
	<meta http-equiv='Content-Type' content='text/html; charset=utf-8' />
	<META HTTP-EQUIV='CACHE-CONTROL' CONTENT='NO-CACHE'>
	<meta name='rating' content='general'>
	<link rel='stylesheet' href='/menu.css' type='text/css' media='screen' />
	<link href='/rss/index.php' rel='alternate' type='application/rss+xml' title='News' />
	<title>Library Genesis</title>
	<!--[if IE 6]>
	<style>
		body {behavior: url('/csshover3.htc');}
		#menu li .drop {background:url('img/drop.gif') no-repeat right 8px; 
	</style>
	<![endif]-->
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-145683333-1"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-145683333-1');
</script><script src="../clipboard.min.js"></script><script src="../jquery-latest.min.js"></script></head><body><script>
	$(document).ready(function() {
		function appendMessage(argument) {
			var div = $('<div>').attr('id', 'message').text('Ad block is installed and active. Please support us by disabling it.');
			var add = $('span').before(div);
		}
		setTimeout(function(){
			if($("ins").css('display') == "none") {
				appendMessage();
			}
		}, 500);
	});
	</script><span></span><table id=main  align="center" border=0>
	<tr>

	<td align="left" valign="top" bgcolor="#F5F6CE" width=200 nowrap></td>
	<td align="center" valign="top" bgcolor="#A9F5BC"><a href="http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=1R6YL7A6N8NVWLB0"><h2>GET</h2></a></td>
	<td align="left" valign="top" bgcolor="#F5F6CE" width=450></td>
	</tr>
	<tr>

	<td bgcolor="#F5F6CE" valign=top><script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- genvert1 -->
<ins class="adsbygoogle"
     style="display:inline-block;width:300px;height:600px"
     data-ad-client="ca-pub-7395443848370277"
     data-ad-slot="7601779340"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script></td>
	<td><a href='/book/index.php?md5=2f2dba2a621b693bb95601c16ed680f8&oftorrent='>Download via torrent </a><input id="textarea-example" value="(Ablex Series in Artificial Intelligence) Larry J. Crockett-The Turing Test and the Frame Problem_ AI's Mistaken Understanding of Intelligence-Ablex Publishing Corporation (1994).gz" type="text" size="9"><button class="btn-clipboard" data-clipboard-target="#textarea-example"> (need rename file)</button><script>new Clipboard(".btn-clipboard");</script><table><tr><td><img src='/img/blank.png' border=0 width=240 style='padding: 5px'></td><td></td></tr><tr><td>Title: The Turing Test and the Frame Problem: AI&#039;s Mistaken Understanding of Intelligence<br>
Author(s): Larry J. Crockett<br>
Publisher: Ablex Publishing Corporation<br>
Year: 1994<br>
ISBN: 9780893919269,0893919268<br></td><td><textarea rows='13' name='bibtext' id='bibtext' readonly cols='40'>@book{book:{643},
   title =     {The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence},
   author =    {Larry J. Crockett},
   publisher = {Ablex Publishing Corporation},
   isbn =      {9780893919269,0893919268},
   year =      {1994},
   series =    {Ablex Series in Artificial Intelligence},
   edition =   {},
   volume =    {},
   url =       {http://gen.lib.rus.ec/book/index.php?md5=2f2dba2a621b693bb95601c16ed680f8}}</textarea></td></tr><tr><td colspan=2><p style='text-align:center'>
<a href='https://www.worldcat.org/search?qt=worldcat_org_bks&q=The%20Turing%20Test%20and%20the%20Frame%20Problem%3A%20AI%27s%20Mistaken%20Understanding%20of%20Intelligence&fq=dt%3Abks'>Search in WorldCat</a> 
<a href='https://www.goodreads.com/search?utf8=✓&query=The%20Turing%20Test%20and%20the%20Frame%20Problem%3A%20AI%27s%20Mistaken%20Understanding%20of%20Intelligence'>Search in Goodreads</a><br>
<a href='https://www.abebooks.com/servlet/SearchResults?tn=The%20Turing%20Test%20and%20the%20Frame%20Problem%3A%20AI%27s%20Mistaken%20Understanding%20of%20Intelligence&pt=book&cm_sp=pan-_-srp-_-ptbook'>Search in AbeBooks</a></td></tr></table></td>
	<td bgcolor="#F5F6CE" valign=top><script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
<!-- genvert2 -->
<ins class="adsbygoogle"
     style="display:inline-block;width:450px;height:600px"
     data-ad-client="ca-pub-7395443848370277"
     data-ad-slot="3816794618"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>


</td>
	</tr>
	<tr>
	<td></td>
	<td></td>
	<td></td>
	
	</tr>
	<tr><td></td><td colspan=2>Both the Turing test and the frame problem have been significant items of discussion since the 1970s in the philosophy of artificial intelligence (AI) and the philisophy of mind. However, there has been little effort during that time to distill how the frame problem bears on the Turing test. If it proves not to be solvable, then not only will the test not be passed, but it will call into question the assumption of classical AI that intelligence is the manipluation of formal constituens under the control of a program. This text explains why there has been less progress in artificial intelligence research than AI proponents would have believed in the mid-1970s. As a first pass, the difficulty of the frame problem would account for some of the lack of progress. An alternative interpretation is that the research paradigm itself is destined to be less productive than might have been hoped. In general termns, the view advanced here is that the future of AI depends on whether the frame problem eventually falls to computational techniques. If it turns out that the frame problem is computationally irreducible, of there is no way to solve it computationally by means of a program operating on formally defined constituents, then an increasing number of experts in the field will reach the conclusion that AI embodies a fundamental misunderstanding of intelligence.</td></tr>
	</table></body></html>`)
	if results == "" {
		t.Error("empty result")
	}
	if !strings.Contains(results, "http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=") {
		t.Error("incorrect URL returned")
	}
}
