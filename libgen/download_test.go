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
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"2F2DBA2A621B693BB95601C16ED680F8"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getBooksdlDownloadURL(book[0]); err != nil {
		t.Error(err)
	}
	if err := DownloadBook(book[0], ""); err != nil {
		t.Error(err)
	}
}

func TestGetDownloadURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"2F2DBA2A621B693BB95601C16ED680F8"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := GetDownloadURL(book[0]); err != nil {
		t.Error(err)
	}
	if book[0].DownloadURL == "" {
		t.Error("download URL empty")
	}
}

func TestGetBokDownloadURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"2F2DBA2A621B693BB95601C16ED680F8"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getBokDownloadURL(book[0]); err != nil {
		if err.Error() != "download limit reached for b-ok.cc" {
			t.Error(err)
		}
	}

	if book[0].DownloadURL == "" {
		t.Error("no valid url found")
	}
	if !strings.Contains(book[0].DownloadURL, "https://b-ok.cc/dl/436993/") {
		t.Errorf("got: %s, expected: https://b-ok.cc/dl/436993/", book[0].DownloadURL)
	}
}

func TestGetBooksdlDownloadURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"2F2DBA2A621B693BB95601C16ED680F8"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getBooksdlDownloadURL(book[0]); err != nil {
		t.Error(err)
	}

	if book[0].DownloadURL == "" {
		t.Error("no valid url found")
	}
	if !strings.Contains(book[0].DownloadURL, "http://80.82.78.13/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=") {
		t.Errorf("got: %s, expected: http://80.82.78.13/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=", book[0].DownloadURL)
	}
}

func TestGetNineThreeURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"2F2DBA2A621B693BB95601C16ED680F8"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getNineThreeURL(book[0]); err != nil {
		t.Error(err)
	}

	if book[0].DownloadURL == "" {
		t.Error("no valid url found")
	}
	if !strings.Contains(book[0].DownloadURL, "http://93.174.95.29") {
		t.Errorf("got: %s, expected: http://93.174.95.29", book[0].DownloadURL)
	}
}

func TestGetHref(t *testing.T) {
	results := findMatch(booksdlReg, []byte(`<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'>
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
	<td align="center" valign="top" bgcolor="#A9F5BC"><a href="http://80.82.78.13/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=1R6YL7A6N8NVWLB0&mirr=1"><h2>GET</h2></a></td>
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
<a href='https://www.goodreads.com/search?utf8=✓&Query=The%20Turing%20Test%20and%20the%20Frame%20Problem%3A%20AI%27s%20Mistaken%20Understanding%20of%20Intelligence'>Search in Goodreads</a><br>
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
	</table></body></html>`))
	if results == nil {
		t.Error("empty result")
	}
	if !strings.Contains(string(results), "http://80.82.78.13/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=") {
		t.Error("incorrect DownloadURL returned")
	}
}

func TestCheckBokDownloadLimit(t *testing.T) {
	results := findMatch(bokDownloadLimit, []byte(`<!DOCTYPE html>
<html>
    <head>
        <title>Downloading The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence</title>
<base href="/">

                        <meta charset="utf-8">		                       
                        <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
                        <meta http-equiv="X-UA-Compatible" content="IE=edge">
                        <meta name="viewport" content="width=device-width, initial-scale=1">
                        <meta name="title" content="Downloading The Turing Test and the Frame Problem: AI&#039;s Mistaken Understanding of Intelligence">
			<meta name="description" content="Downloading The Turing Test and the Frame Problem: AI&#039;s Mistaken Understanding of Intelligence | B–OK. Download books for free. Find books">
			<meta name="robots" content="index,all">
			<meta name="distribution" content="global">
			<meta http-equiv="cache-control" content="no-cache">
			<meta http-equiv="pragma" content="no-cache">

                        <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
                        <link rel="icon" type="image/png" href="/favicon-32x32.png" sizes="32x32">
                        <link rel="icon" type="image/png" href="/favicon-16x16.png" sizes="16x16">
                        <link rel="manifest" href="/manifest.json">
                        <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5">
                        <meta name="apple-mobile-web-app-title" content="Z-Library">
                        <meta name="application-name" content="Z-Library">
                        <meta name="theme-color" content="#ffffff">

                        <meta name="propeller" content="49c350d528ba144cace841cac74260ab">
	
<!-- CSS SET -->
<!-- JS SET --> 
<link REL="SHORTCUT ICON" HREF="/favicon.ico">
        <link rel="search" type="application/opensearchdescription+xml" href="/search.xml" title="Search for books in the library B-OK.org" />

                    <link rel="stylesheet" type="text/css" href="/resources/build/global.css?0.188" />
            <script type="text/javascript" src="/resources/build/global.js?0.188"></script>
        
                
        <script>
            const CurrentUser = new User(null)

                    </script>
    </head>

    <body style="margin:0px;padding:0px;" class="download">
        <table border="0" height="100%" width="100%" style="height:100%;" cellpadding="0" cellspacing="0">
            <tbody>
                <tr style="height:10px;">
                    <td>
                        <div class="container-fluid">
                            <style>
    div#colorBoxes ul li.active:nth-child(1)
    {
        border: 1px solid #378096;
        box-shadow: 0 0 6px #7DBCCF;
        border-top: 0px;

    }

    div#colorBoxes ul li.active:nth-child(2)
    {
        border: 1px solid #6e9b41;
        box-shadow: 0 0 6px #a4e861;
        border-top: 0px;
    }
</style>

<div class="row">
    <div class="col-md-12">
        <div id="colorBoxes" class="darkShadow">
            <ul>
                <li style="background: #49afd0;" class="active">
                    <a href="/">
                        <span class="hidden-xs">5,043,823</span>
                        Books                    </a>
                </li>

                <li style="background: #8ecd51;" class="">
                    <a href="http://booksc.xyz/">
                        <span class="hidden-xs">77,511,848</span>
                        Articles                    </a>
                </li>

                <li style="background: #90a5a8;" class="hidden-xs"><a href="https://z-lib.org">ZLibrary Home</a></li>
                <li style="background: #90a5a8;" class="visible-xs"><a href="https://z-lib.org">Home</a></li>
            </ul>
        </div>

        <div role="navigation" class="navbar-default navbar-right" style="background-color: transparent;">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false" onclick="setTimeout(function () {$('#main-menu-dropdown').click()}, 100)">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
            </div>

            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1" style="clear: both;">
                <ul class="nav navbar-nav navbar-right" style="">
                                            <li class="dropdown">
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Sign in <span class="caret"></span></a>
                            <ul class="dropdown-menu">
                                <li><a href="http://singlelogin.org?from=b-ok.cc">Login</a></li>
                                <li><a href="https://singlelogin.org/registration.php">Registration</a></li>
                            </ul>
                        </li>
                    
                    
                    <li>
                    <a href="/howtodonate.php" style="color:#8EB46A;">Donate</a>                    </li>
                    <li class="dropdown">
                        <a href="#" id="main-menu-dropdown" class="dropdown-toggle hidden-xs" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="true">
                            <span style="font-size: 120%;" class="glyphicon glyphicon-menu-hamburger" aria-hidden="true"></span>
                                                    </a>
                        <ul class="dropdown-menu">
                                                        <li class="dropdown-header">Books</li>
                                                            <li><a href="/book-add.php">Add book</a></li>
                                <li><a href="/categories">Categories</a></li>
                                <li><a href="/popular.php">Most Popular</a></li>
                                                                                    <li><a href="/recently.php">Recently Added</a></li>
                                                        <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Z-Library Project</li>
                            <li><a href="/top-zlibrarians.php">Top Z-Librarians</a></li>
                            <li><a href="/blog/">Blog</a></li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</div>

                        </div>
                    </td>
                </tr>
                <tr>
                    <td>
                        <div class="container" style="max-width: 1000px;padding-bottom: 40px;">
                            <div class="row">
                                <div class="col-md-12 itemFullText">
                                    <div class="row">
    <div class="col-md-12">
        <div style="text-align: center; font-size: 18pt; color: #49AFD0; line-height: 160%; margin:50px 0;">
            WARNING: There are more than 5 downloads from your IP <span style="color:#8ECD51;">68.235.43.101</span> during last 24 hours. <br/><br/> Please, <a href="http://singlelogin.org?from=b-ok.cc">login</a> into your account or  <a href="https://b-ok.cc/registration.php">complete a simple registration</a> (do not create more than one account) to download more books.        </div>
    </div>
</div>

    <script type="text/javascript">
        CurrentUser.deleteDownloadedBook(436993)
    </script>

<!-- https://bootsnipp.com/snippets/a68Do -->

<div class="table-responsive">
    <div class="membership-pricing-table" style="margin-top:40px;">
        <table class="table table-responsive" style="width:98%;">
            <tbody>
                <tr>
                    <th>&nbsp;</th>
                    <th class="plan-header plan-header-free">
                        <div class="header-plan-inner">
                            <span class="recommended-plan-ribbon">Your current status</span>                            <div class="pricing-plan-name">GUEST</div>
                            <div class="pricing-plan-period">FREE</div>
                        </div>
                    </th>
                    <th class="plan-header plan-header-blue">
                        <div class="header-plan-inner">
                                                        <div class="pricing-plan-name">BASIC MEMBER</div>
                            <div class="pricing-plan-period">FREE</div>
                        </div>
                    </th>
                    <th class="plan-header plan-header-standard">
                        <div class="header-plan-inner">
                                                        <div class="pricing-plan-name">PREMIUM</div>
                            <div class="pricing-plan-period">Contributors</div>
                        </div>
                    </th>
                </tr>
                <tr>
                    <td>Downloads (daily)</td>
                    <td>5</td>
                    <td>10</td>
                    <td>&gt; 10 up to 999<br/><small>depends on contribution</small></td>
                </tr>
                <tr>
                    <td>Download speed</td>
                    <td>up to 1 MBps</td>
                    <td>up to 1 MBps</td>
                    <td>Unrestricted</td>
                </tr>
                <tr>
                    <td>Downloads history</td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td>Send books to email</td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td>Personal recommendations</td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td><a href="/info/howtokindle.php" target="_blank">Send-to-Kindle feature</a></td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td>Files converter<br/><span style="font-size: 80%;">Convert books to epub, pdf, mobi, fb2 and other formats </span></td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td>No ADs</td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td>Access to book4you.org</td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-no glyphicon glyphicon-remove"></span></td>
                    <td><span class="icon-yes glyphicon glyphicon-ok"></span></td>
                </tr>
                <tr>
                    <td>Price</td>
                    <td>FREE</td>
                    <td>FREE</td>
                    <td>Donate any amount</td>
                </tr>
                <tr>
                    <td>How to proceed</td>
                    <td><span style="color: #209E61">Relax. You're already here</span></td>
                    <td><a href="http://singlelogin.org?from=b-ok.cc">Sign In</a> or <a href="http://singlelogin.org/registration.php">Sign Up</a></td>
                    <td><a href="/howtodonate.php" class="btn btn-info"  role="button">Make a donation</a></td>
                </tr>
            </tbody></table>
    </div>
</div>

<style>

.membership-pricing-table table, .membership-pricing-table table td {
    border: 1px solid #ebebeb;
}

    .membership-pricing-table table .icon-no,.membership-pricing-table table .icon-yes {
        font-size: 22px
    }

    .membership-pricing-table table .icon-no {
        color: #a93717
    }

    .membership-pricing-table table .icon-yes {
        color: #209e61
    }

    .membership-pricing-table table .plan-header {
        text-align: center;
        font-size: 48px;
        border: 1px solid #e2e2e2;
        padding: 24px 0;
    }

    .membership-pricing-table table .plan-header-free {
        background-color: #eee;
        color: #555
    }

    .membership-pricing-table table .plan-header-blue {
        color: #fff;
        background-color: #61a1d1;
        border-color: #3989c6
    }

    .membership-pricing-table table .plan-header-standard {
        color: #fff;
        background-color: #ff9317;
        border-color: #e37900
    }

    .membership-pricing-table table td {
        text-align: center;
        width: 25%;
        padding: 7px 10px;
        background-color: #fff;
        font-size: 14px;
        -webkit-box-shadow: 0 1px 0 #fff inset;
        box-shadow: 0 1px 0 #fff inset
    }

    .membership-pricing-table table td {
        border: 1px solid #ebebeb;
        vertical-align: middle !important;
    }

    .membership-pricing-table table tr td:first-child {
        background-color: transparent;
        text-align: right;
        width: 24%;
    }

    .membership-pricing-table table tr td:nth-child(5) {
        background-color: #FFF
    }

    .membership-pricing-table table tr:first-child td,.membership-pricing-table table tr:nth-child(2) td {
        -webkit-box-shadow: none;
        box-shadow: none;
        vertical-align: middle;
    }

    .membership-pricing-table table tr:first-child th:first-child {
        /*border-top-color: transparent;
        border-left-color: transparent;
        border-right-color: #e2e2e2*/
    }

    .membership-pricing-table table tr:first-child th .pricing-plan-name {
        font-size: 22px
    }

    .membership-pricing-table table tr:first-child th .pricing-plan-period {
        margin-top: 7px;
        font-size: 25%
    }

    .membership-pricing-table table .header-plan-inner {
        position: relative
    }

    .membership-pricing-table table .plan-head {
        box-sizing: content-box;
        background-color: #ff9c00;
        border: 1px solid #cf7300;
        position: absolute;
        top: -33px;
        left: -1px;
        height: 30px;
        width: 100%;
        border-bottom: none
    }

    .membership-pricing-table table .recommended-plan-ribbon {
        white-space: nowrap;
        box-sizing: content-box;
        color: #dc3b5d;
        position: absolute;
        padding: 3px 6px;
        font-size: 12px!important;
        font-weight: 600;
        left: 0px;
        top: -50px;
        z-index: 99;
        width: 100%;
    }

    .membership-pricing-table table tr td:first-child{
        color: #555;
    }

</style>                                                                    </div>
                            </div>
                        </div>
                    </td>
                </tr>
                <tr style="height:60px">
                    <td id="footer" valign="top">
                        <div class="container-fluid">
<!-- footer begin -->
<div class="row">
    <div class="col-sm-8 footer-copyright">
        <i>
            Free ebooks since 2009.
            <a class="footer-mailto" href="mailto: support@bookmail.org">
                support@bookmail.org            </a>
        </i>

        <span class="hidden-xs" style="margin:0 0 0 15px;"> <a href="/faq.php">FAQ</a></span>
        <span class="hidden-xs" style="margin:0 0 0 15px;"> <a href="/blog/">Blog</a></span>
    </div>

    <div class="col-sm-4">
        <div class="pull-right footer-nav-right" role="navigation">
            <ul class="nav navbar-nav">
                <li class="visible-xs-block"><a href="/faq.php">FAQ</a></li>
                <li class="visible-xs-block"><a href="/blog/">Blog</a></li>

                <li><a href="/privacy.php">Privacy</a></li>
                <li><a href="/dmca.php">DMCA</a></li>
                <li class="dropup">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">English <span class="caret"></span></a>
                    <ul class="dropdown-menu">
                        <li><a onclick="setLanguage('en'); return false;" href="//en.b-ok.cc/dl/436993/1f4c46">English</a></li><li><a onclick="setLanguage('ru'); return false;" href="//ru.b-ok.cc/dl/436993/1f4c46">Русский</a></li><li><a onclick="setLanguage('ua'); return false;" href="//ua.b-ok.cc/dl/436993/1f4c46">Українська</a></li><li><a onclick="setLanguage('pl'); return false;" href="//pl.b-ok.cc/dl/436993/1f4c46">Polski</a></li><li><a onclick="setLanguage('it'); return false;" href="//it.b-ok.cc/dl/436993/1f4c46">Italiano</a></li><li><a onclick="setLanguage('es'); return false;" href="//es.b-ok.cc/dl/436993/1f4c46">Español</a></li><li><a onclick="setLanguage('zh'); return false;" href="//zh.b-ok.cc/dl/436993/1f4c46">汉语</a></li><li><a onclick="setLanguage('id'); return false;" href="//id.b-ok.cc/dl/436993/1f4c46">Bahasa Indonesia</a></li><li><a onclick="setLanguage('in'); return false;" href="//in.b-ok.cc/dl/436993/1f4c46">हिन्दी</a></li><li><a onclick="setLanguage('pt'); return false;" href="//pt.b-ok.cc/dl/436993/1f4c46">Português</a></li><li><a onclick="setLanguage('jp'); return false;" href="//jp.b-ok.cc/dl/436993/1f4c46">日本語</a></li><li><a onclick="setLanguage('de'); return false;" href="//de.b-ok.cc/dl/436993/1f4c46">Deutsch</a></li><li><a onclick="setLanguage('fr'); return false;" href="//fr.b-ok.cc/dl/436993/1f4c46">Français</a></li><li><a onclick="setLanguage('th'); return false;" href="//th.b-ok.cc/dl/436993/1f4c46">ภาษาไทย</a></li><li><a onclick="setLanguage('el'); return false;" href="//el.b-ok.cc/dl/436993/1f4c46">ελληνικά </a></li><li><a onclick="setLanguage('ar'); return false;" href="//ar.b-ok.cc/dl/436993/1f4c46">اللغة العربية</a></li>                    </ul>
                </li>
            </ul>
        </div>
    </div>
</div></div>
                    </td>
                </tr>
            </tbody>
        </table>

        <!-- ggAdditionalHtml -->
        
    <script>
        var Config = {"currentLanguage":"en","L":{"89":"rub.","90":"The file is located on an external resource","91":"It is a folder","92":"File from disk storage","93":"File is aviable by direct link","94":"Popular","95":"Limitation of downloading: no more than 2 files at same time","96":"Size","97":" Language","98":"Category"}};
    </script>
    <!--LiveInternet counter--><script type="text/javascript">
new Image().src = "//counter.yadro.ru/hit;bookzz?r"+
escape(document.referrer)+((typeof(screen)=="undefined")?"":
";s"+screen.width+"*"+screen.height+"*"+(screen.colorDepth?
screen.colorDepth:screen.pixelDepth))+";u"+escape(document.URL)+
";"+Math.random();</script><!--/LiveInternet-->

<iframe name="uploader" id="uploader" style="border:0px solid #ddd; width:90%; display:none;"></iframe>        <!-- /ggAdditionalHtml -->
        
                
        <script>
            if (typeof pagerOptions !== "undefined" && pagerOptions) {
                $('div.paginator').paginator(pagerOptions);
            }
        </script>
    </body>
</html>`))
	if results == nil {
		t.Error("empty result")
	}

	if !strings.Contains(string(results), "WARNING: There are more than 5 downloads from your IP") {
		t.Error("incorrect DownloadURL returned")
	}
}
