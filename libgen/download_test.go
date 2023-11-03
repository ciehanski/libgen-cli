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
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getLibraryLolURL(book[0], false); err != nil {
		t.Error(err)
	}
	if err := DownloadBook(book[0], ""); err != nil {
		t.Error(err)
	}
}

func TestGetDownloadURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := GetDownloadURL(book[0], false); err != nil {
		t.Error(err)
	}
	if book[0].DownloadURL == "" {
		t.Error("download URL empty")
	}
}

func TestLibgenPMDownloadURL(t *testing.T) {
	t.Skipf("Skipping, does not pass in GitHub Actions")
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getLibgenPMURL(book[0]); err != nil {
		t.Error(err)
	}

	if book[0].DownloadURL == "" {
		t.Error("no valid url found")
	}
	if !strings.Contains(book[0].DownloadURL, "https://libgen.rocks/get.php?md5=1794743bb21d72736ffe64d66dca9f0e&key=") {
		t.Errorf("got: %s, expected: https://libgen.rocks/get.php?md5=1794743BB21D72736FFE64D66DCA9F0E&key=", book[0].DownloadURL)
	}
}

func TestGetLibraryLolURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getLibraryLolURL(book[0], false); err != nil {
		t.Error(err)
	}

	if book[0].DownloadURL == "" {
		t.Error("no valid url found")
	}
	if !strings.Contains(book[0].DownloadURL, "https://download.library.lol/main/1440000/1794743bb21d72736ffe64d66dca9f0e") {
		t.Errorf(`got: %s, expected: https://download.library.lol/main/1440000/1794743bb21d72736ffe64d66dca9f0e`, book[0].DownloadURL)
	}
}

func TestGetHref(t *testing.T) {
	results := findMatch(libgenPMReg, []byte(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<META HTTP-EQUIV="CACHE-CONTROL" CONTENT="max-age=72000, must-revalidate">
	<meta name="rating" content="general">
	<!--<link href="/rss/index.php" rel="alternate" type="application/rss+xml" title="News" />-->
	<link rel="shortcut icon" href="/img/favicon.ico" type="image/x-icon">
	<title>Library Genesis</title>
		
	<!--[if IE 6]>
	<style>
		body {behavior: url("/csshover3.htc");}
		#menu li .drop {background:url("img/drop.gif") no-repeat right 8px; 
	</style>
	<![endif]-->
<link rel="stylesheet" href="../css/bootstrap.min.css">	
	
<link href="/css/font.min.css" rel="stylesheet">	
<style>
nav.navbar .dropdown:hover > .dropdown-menu {
 display: block; 
}
.bd-placeholder-img {
	font-size: 1.125rem;
	text-anchor: middle;
	-webkit-user-select: none;
	-moz-user-select: none;
	-ms-user-select: none;
	user-select: none;
}
@media (min-width: 768px) {
			.bd-placeholder-img-lg {
			font-size: 3.5rem;
		}
	}

.panel-heading .accordion-toggle:after {
    font-family: "Glyphicons Halflings";  
    content: "\e114";    
    float: right;       
    color: grey;         
}
.panel-heading .accordion-toggle.collapsed:after {
    content: "\e080";   
}
.tooltip-inner {
    max-width: 350px;
    width: 350px; 
}
h1 {
	display: block; 
	font-size: 1.8rem; 
	font-weight: bold; 
	font-family: Georgia, "Times New Roman", Times, serif;  color: #A00000; 
}
#tablelibgen td { 
	font-family: "Pt Sans", Tahoma, Helvetica, sans-serif; 
	margin: 0; 
	padding: 0em 3px; 
	font-size: 1rem;
}

#tablelibgen1 td { 
	font-family: "Pt Sans", Tahoma, Helvetica, sans-serif; 
	margin: 0; 
	padding: 0em 3px; 
	font-size: 1rem;
}

.taghide {
    display: none; 
}
.taghide + label ~ div {
    display: none;
}
/* оформляем текст label */
.taghide + label {
    display: inline-block; 
}
/* вид текста label при активном переключателе */

/* когда чекбокс активен показываем блоки с содержанием  */
.taghide:checked + label + div {
    display: block; 
}



/*.navbar {
	background-color: #BBBBBB;
}*/
	</style>

	<link rel="stylesheet" href="/css/dark-mode.css">
	<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
<style>p {margin: 0;}</style>
</head>
<body><script data-ad-client="ca-pub-4139850031026202" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>    
<nav class="navbar navbar-expand-md navbar-dark bg-secondary  mb-4">
  
   <a class="navbar-brand" href="/index.php">
    <img src="/img/logo.png"  height="30" alt="">
  </a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarCollapse">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item active">
        <a class="nav-link" href="/community/app.php/article/news">NEWS <span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item active">
        <a class="nav-link" href="/community/">FORUM <span class="sr-only">(current)</span></a>
      </li>
	
      <li class="nav-item dropdown">
<a class="btn btn-secondary dropdown-toggle" href="/community/ucp.php?mode=login" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          LOGIN
        </a>
        <div class="dropdown-menu" aria-labelledby="dropdown01">    
          <a class="dropdown-item" href="/community/ucp.php?mode=register">Register</a>
        </div>
      </li>
      <li class="nav-item dropdown">
        <a class="btn btn-secondary dropdown-toggle" href="#" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          DOWNLOAD
        </a>
        <div class="dropdown-menu" aria-labelledby="dropdown01">      

          <a class="dropdown-item" href="/mirrors.php">Mirrors</a>
          <a class="dropdown-item" href="http://libgenfrialc7tguyjywa36vtrdcplwpxaw43h6o63dmmwhvavo5rqqd.onion/">TOR</a>

	<div class="dropdown-divider"></div>
         <h6 class="dropdown-header">P2P</h6>
          <a class="dropdown-item" href="/torrents/">Torrents</a>
          <a class="dropdown-item" href="/nzb/">Usenet (*.nzb)</a>
          <a class="dropdown-item" href="https://phillm.net/libgen-stats-table.php">Torrents status</a>




	<div class="dropdown-divider"></div>
         <h6 class="dropdown-header">DB Dumps</h6>
          <a class="dropdown-item" href="/dirlist.php?dir=dbdumps">Libgen</a>
          <a class="dropdown-item" href="http://libgen.rs/dbdumps/">libgen.rs (gen.lib.rus.ec)</a>

	<div class="dropdown-divider"></div>
 	 <a class="dropdown-item" href="/comics0/">Unsorted comics</a>
 	 <a class="dropdown-item" href="/magz0/">Unsorted magz</a>
 	 <a class="dropdown-item" href="/fict0/">Unsorted fiction</a>
        </div>

      </li>

      <li class="nav-item dropdown">
        <a class="btn btn-secondary dropdown-toggle" href="librarian.php" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          UPLOAD
        </a>
        <div class="dropdown-menu" aria-labelledby="dropdown01">  
          <a class="dropdown-item" href="ftp://ftp.libgen.lc/upload/">FTP</a> 
        </div>
      </li>

      <li class="nav-item dropdown">
        <a class="btn btn-secondary dropdown-toggle" href="/index.php?req=fmode:last&topics1=all" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          LAST
        </a>

        <div class="dropdown-menu" aria-labelledby="dropdown01">
	<a class="dropdown-item" href="/index.php?req=fmode:last&topics1=all"><b>Files</b></a>

          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=l">Libgen</a>
          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=a">Scientific Articles</a> 
          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=f">Fiction</a> 
          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=c">Comics</a> 
          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=m">Magazines</a> 
          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=s">Standards</a> 
          <a class="dropdown-item" href="/index.php?req=fmode:last&topics%5B%5D=r">Fiction RUS</a>
	<div class="dropdown-divider"></div>
          <a class="dropdown-item" href="/index.php?req=mode:last&curtab=e">Editions</a> 
          <a class="dropdown-item" href="/index.php?req=mode:last&curtab=s">Series</a>
          <a class="dropdown-item" href="/index.php?req=mode:last&curtab=p">Publishers</a> 
        <!--  <a class="dropdown-item" href="/index.php?req=mode:last&curtab=f">Files</a> -->
          <a class="dropdown-item" href="/index.php?req=mode:last&curtab=a">Authors</a> 
          <a class="dropdown-item" href="/index.php?req=mode:last&curtab=w">Works</a>


  
        </div>


      </li>

      <li class="nav-item dropdown">
        <a class="btn btn-secondary dropdown-toggle" href="#" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          OTHERS
        </a>

        <div class="dropdown-menu" aria-labelledby="dropdown01">  
          <a class="dropdown-item" href="json.php">API</a> 
          <a class="dropdown-item" href="top.php">Top 100 users</a> 
          <a class="dropdown-item" href="stat.php">Stats</a>
          <a class="dropdown-item" href="batchsearchindex.php">Batch search</a>  
          <a class="dropdown-item" href="biblioservice.php">Bibliographic services</a>
          <a class="dropdown-item" href="http://libruslib.ucoz.com/index/libgen_bibliotekar/0-5">Libgen librarian for desktop</a>


          <a class="dropdown-item" href="/code/">Source (PHP)</a>
          <a class="dropdown-item" href="/soft/">LG soft</a>
          <!--<a class="dropdown-item" href="/import/">Import local files in LG format</a>-->
          <a class="dropdown-item" href="https://b-ok.cc/fulltext/">Full text search</a>



        </div>
      </li>



      <li class="nav-item dropdown">
        <a class="btn btn-secondary dropdown-toggle" href="topics.php" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          TOPICS
        </a>
      </li>


      <li class="nav-item dropdown">
        <a class="btn btn-secondary dropdown-toggle" href="#" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          LINKS
        </a>

        <div class="dropdown-menu" aria-labelledby="dropdown01">  


          
          <a class="dropdown-item" href="http://sci-hub.ru">Sci-hub</a> 
          <a class="dropdown-item" href="http://magzdb.org">Magzdb.org</a>

          <a class="dropdown-item" href="http://nlr.ru/rlin/Periodika_rus.php">РНБ</a>
          <a class="dropdown-item" href="http://rsl.ru/">РГБ</a>
          <a class="dropdown-item" href="https://loc.gov/">LOC</a>
          <a class="dropdown-item" href="https://comicvine.gamespot.com/">ComicVine</a>
          <a class="dropdown-item" href="https://cyberleninka.ru/">Cyberleninka</a>
          <a class="dropdown-item" href="https://lib.rus.ec/">Lib.rus.ec</a>
          <a class="dropdown-item" href="http://flibusta.net/">Flibusta.net</a>
          <a class="dropdown-item" href="https://goodreads.com/">Goodreads.com</a>
          <a class="dropdown-item" href="https://worldcat.org/">Worldcat.org</a>
          <a class="dropdown-item" href="https://wiki.archiveteam.org/">Archive team</a>
          <a class="dropdown-item" href="https://www.reddit.com/r/libgen/">Reddit</a>

        </div>

      </li>


      <li class="nav-item dropdown">
        <a class="btn btn-secondary" href="index.php?req=mode:req&curtab=e" role="button" id="dropdownMenuLink"  aria-haspopup="true" aria-expanded="false">
          WANTED
        </a>
      </li>

    </ul>
  </div>

  <div class="nav-link">

    <div class="custom-control custom-switch">
      <input type="checkbox" class="custom-control-input" id="darkSwitch">
      <label class="custom-control-label" for="darkSwitch">🌓</label>
    </div>
    <script src="/js/dark-mode-switch.js"></script>
  </div>
   <a class="navbar-brand" href="setlang.php?md5=1794743BB21D72736FFE64D66DCA9F0E&lang=ru">RU</a>
</nav>
<span></span><table id=main  align="center" border=1>
		<tr>
	
		<td align="left" valign="top" bgcolor="#F5F6CE" width=200 nowrap></td>
		<td align="center" valign="top" bgcolor="#A9F5BC"><a href="get.php?md5=1794743bb21d72736ffe64d66dca9f0e&key=WBYEV7R2TZE7NEDZ"><h2>GET</h2></a></td>
		<td align="left" valign="top" bgcolor="#F5F6CE" width=450></td>
		</tr>
		<tr>
	
		<td bgcolor="#F5F6CE" valign=top><script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-4139850031026202"
     crossorigin="anonymous"></script>
<!-- skyscraper1 -->
<ins class="adsbygoogle"
     style="display:block"
     data-ad-client="ca-pub-4139850031026202"
     data-ad-slot="5706997950"
     data-ad-format="auto"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script></td>
		<td><table><tr><td><a href="/covers/1440000/1794743bb21d72736ffe64d66dca9f0e.jpg"><img src="/covers/1440000/1794743bb21d72736ffe64d66dca9f0e.jpg" width=300></a></td><td></td></tr>
<tr><td>Title: Getting Started with Kubernetes<br>
Author(s): Jonathan Baier<br>
Publisher: Packt Publishing<br>
Year: 2015<br>
ISBN: 1784394033; 9781784394035<br></td><td><textarea rows='13' name='bibtext' id='bibtext' readonly cols='40'>@book{book:{92476306},
   title =     {Getting Started with Kubernetes},
   author =    {Jonathan Baier},
   publisher = {Packt Publishing},
   isbn =      {1784394033; 9781784394035},
   year =      {2015},
   url =       {libgen.li/file.php?md5=1794743bb21d72736ffe64d66dca9f0e}}</textarea></td></tr>
<tr><td colspan=2><p style='text-align:center'>
<a href='https://www.worldcat.org/search?qt=worldcat_org_bks&q=Getting%20Started%20with%20Kubernetes&fq=dt%3Abks'>Search in WorldCat</a> 
<a href='https://www.goodreads.com/search?utf8=✓&query=Getting%20Started%20with%20Kubernetes'>Search in Goodreads</a><br>
<a href='https://www.abebooks.com/servlet/SearchResults?tn=Getting%20Started%20with%20Kubernetes&pt=book&cm_sp=pan-_-srp-_-ptbook'>Search in AbeBooks</a></td></tr></table></td>
		<td bgcolor="#F5F6CE" valign=top><script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-4139850031026202"
     crossorigin="anonymous"></script>
<!-- skyscraper3fixed -->
<ins class="adsbygoogle"
     style="display:block"
     data-ad-client="ca-pub-4139850031026202"
     data-ad-slot="2486455165"
     data-ad-format="auto"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script></td>
		</tr>

		<tr><td></td><td colspan=2></td></tr>
		<tr><td colspan=3 bgcolor="#F5F6CE"><script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-4139850031026202"
     crossorigin="anonymous"></script>
<!-- horizont1 -->
<ins class="adsbygoogle"
     style="display:block"
     data-ad-client="ca-pub-4139850031026202"
     data-ad-slot="6979435185"
     data-ad-format="auto"
     data-full-width-responsive="true"></ins>
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script></td></tr>
		</table><nav class="navbar sticky-bottom navbar-expand-sm navbar-dark bg-secondary">
  <div class="collapse navbar-collapse" id="navbarCollapse">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item">
	    <a class="nav-link" href="#" data-toggle="modal" data-target="#dmcamodal">DMCA</a>
      </li>
      <li class="nav-item">
	    <a class="nav-link" href="#" data-toggle="modal" data-target="#aboutmodal">ABOUT</a>
      </li>
      <li class="nav-item">
	    <a class="nav-link" href="#" data-toggle="modal" data-target="#donatemodal" >DONATE</a>
      </li>
	
      <li class="nav-item">
	    <a class="nav-link" href="/gdrp.php">GDRP</a>
      </li>
    </ul>
	<span class="navbar-text">Users online 1873</span>
  </div>
</nav>

<!-- Modal Donate -->
<div class="modal fade text-dark" id="donatemodal" tabindex="-1" aria-labelledby="donatemodalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="donatemodalLabel">Bitcoin</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <a href="bitcoin://1HEUTKLrWggDjrUQtX59rUmpK9ckxEXFJb">1HEUTKLrWggDjrUQtX59rUmpK9ckxEXFJb</a>
      </div>
    </div>
  </div>
</div>

<!-- Modal About -->
<div class="modal fade text-dark" id="aboutmodal" tabindex="-1" aria-labelledby="aboutmodalLabel" aria-hidden="true">
  <div class="modal-dialog modal-lg">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="aboutmodalLabel">About</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">


<div id="about">
The Library Genesis aggregator is a community aiming at collecting and cataloging items descriptions for the most part of scientific, 
scientific and technical directions, as well as file metadata. In addition to the descriptions, 
the aggregator contains only links to third-party resources hosted by users. 
All information posted on the website is collected from publicly available public Internet resources and is intended solely for informational purposes.  
</div>
      </div>
    </div>
  </div>
</div>

<!-- Modal DMCA -->
<div class="modal fade text-dark" id="dmcamodal" tabindex="-1" aria-labelledby="dmcamodalLabel" aria-hidden="true">
  <div class="modal-dialog modal-lg">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="dmcamodalLabel">About</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">

<div id="dmca">
Library Genesis - aggregator items is a website that collects and organizes online items from users. 
Item aggregation is done for fact-finding purposes, and website Library Genesis respects the rights of copyright holders and respect dcma.

     Removing Content From Library Genesis / DMCA Policy
     Library Genesis respects the intellectual property of others.
</div>

    <div class="dmca">
     If you believe that your copyrighted work has been copied in a way that constitutes copyright infringement and is accessible on this site, you may notify our copyright agent, as set forth in the Digital Millennium Copyright Act of 1998 (DMCA). For your complaint to be valid under the DMCA, you must provide the following information when providing notice of the claimed copyright infringement:
</div>
    <div class="dmca">
     * A physical or electronic signature of a person authorized to act on behalf of the copyright owner Identification of the copyrighted work claimed to have been infringed <br />
     * Identification of the material that is claimed to be infringing or to be the subject of the infringing activity and that is to be removed <br />
     * Information reasonably sufficient to permit the service provider to contact the complaining party, such as an address, telephone number, and, if available, an electronic mail address <br />
     * A statement that the complaining party "in good faith believes that use of the material in the manner complained of is not authorized by the copyright owner, its agent, or law" <br />
     * A statement that the "information in the notification is accurate", and "under penalty of perjury, the complaining party is authorized to act on behalf of the owner of an exclusive right that is allegedly infringed" <br />
     The above information must be submitted as a written, faxed or emailed notification to the following Designated Agent: <a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="aec7cfc0d4c2c7cceededcc1dac1c0c3cfc7c280cdc1c380">[email&#160;protected]</a> Appeals will be reviewed within 72 hours.</div>


      </div>
    </div>
  </div>
</div>


	<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script src="/js/popper.min.js"></script>
	<!--<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/2.9.2/umd/popper.min.js sha512-2rNj2KJ+D8s1ceNasTIex6z4HWyOnEYLVC3FigGOmyQCZc2eBXKgOxQmo3oKLHyfcj53uz4QMsRCWNbLd32Q1g==" crossorigin="anonymous"></script>-->
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.min.js" integrity="sha384-w1Q4orYjBQndcko6MimVbzY0tgp4pWB4lZ7lr30WKz0vr/aWKhXdBNmNb5D92v7s" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
	<script src="/js/form-validation.js"></script>
	<script>
$('[data-toggle="tooltip"]').tooltip();
$('.btn-tooltip-bottom').tooltip({
    placement: 'bottom'
});
</script>

	
<script>(function(){var js = "window['__CF$cv$params']={r:'72a6567bbcb4631a',m:'EumujVqP1QjoktrPy_1UUOLv7Hu3Dj1HmdxSXe8Zs3s-1657760598-0-ARfqW1o72+1YZyH2PfpK1I0QNIQcA7KzA3a+Jdh1fb9cwkYKfxCxCc/Bi2Clmu5XheKebECrupStunPYHSeDw9qZpP8SUREw3ZY3eRmzkPhpGEhnXr3FT4XdiTmXbb6s0ZKumIZsrfLPcNhWZMAx5ol77nKvVVz8Vr15XfxWJRTn',s:[0x162d84b1e4,0x5a41316e90],u:'/cdn-cgi/challenge-platform/h/g'};var now=Date.now()/1000,offset=14400,ts=''+(Math.floor(now)-Math.floor(now%offset)),_cpo=document.createElement('script');_cpo.nonce='',_cpo.src='/cdn-cgi/challenge-platform/h/g/scripts/alpha/invisible.js?ts='+ts,document.getElementsByTagName('head')[0].appendChild(_cpo);";var _0xh = document.createElement('iframe');_0xh.height = 1;_0xh.width = 1;_0xh.style.border = 'none';_0xh.style.visibility = 'hidden';document.body.appendChild(_0xh);function handler() {var _0xi = _0xh.contentDocument || _0xh.contentWindow.document;if (_0xi) {var _0xj = _0xi.createElement('script');_0xj.innerHTML = js;_0xi.getElementsByTagName('head')[0].appendChild(_0xj);}}if (document.readyState !== 'loading') {handler();} else if (window.addEventListener) {document.addEventListener('DOMContentLoaded', handler);} else {var prev = document.onreadystatechange || function () {};document.onreadystatechange = function (e) {prev(e);if (document.readyState !== 'loading') {document.onreadystatechange = prev;handler();}};}})();</script></body>
</html>
    `))
	if results == nil {
		t.Error("empty result")
	}
	if !strings.Contains(string(results), "get.php?md5=1794743bb21d72736ffe64d66dca9f0e&key=") {
		t.Errorf("incorrect DownloadURL returned. got %s", string(results))
	}
}
