// Copyright © 2019 Ryan Ciehanski <ryan@ciehanski.com>
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
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	results, err := Search(
		"test",
		1,
		false,
		false,
		"",
	)
	if err != nil {
		t.Error(err)
	}
	if strings.ToUpper(results[0].Md5) != "2F2DBA2A621B693BB95601C16ED680F8" {
		t.Errorf("got: %s, expected: 2F2DBA2A621B693BB95601C16ED680F8", strings.ToUpper(results[0].Md5))
	}
}

func TestGetDetails(t *testing.T) {
	books, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8", "06E6135019C8F2F43158ABA9ABDC610E"},
		GetWorkingMirror(SearchMirrors),
		false,
		false,
		"")
	if err != nil {
		t.Error(err)
	}
	if books[0].Title != "The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence" {
		t.Error()
	}
	if books[1].Title != "You failed your math test, Comrade Einstein (about Soviet antisemitism)" {
		t.Error()
	}
}

func TestParseHashes(t *testing.T) {
	response := `<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'>
<html xmlns='http://www.w3.org/1999/xhtml'>
<head>
    <meta http-equiv='Content-Type' content='text/html; charset=utf-8' />
    <META HTTP-EQUIV='CACHE-CONTROL' CONTENT='NO-CACHE'>
    <meta name='robots' content='noindex,nofollow'>
    <meta name='description' content='Library Genesis is a scientific community targeting collection of books on natural science disciplines and engineering.'>
    <meta name='rating' content='general'>

    <link rel='stylesheet' href='../menu.css' type='text/css' media='screen' />
    <title>Library Genesis</title>
    <!--[if IE 6]>
    <style>
        body {behavior: url('../csshover3.htc');}
        #menu li .drop {background:url('img/drop.gif') no-repeat right 8px;
    </style>
    <![endif]-->
</head><body>
<ul id="menu">
    <li><a href="../setlang.php?lang=ru">RU</a><!-- Начало пункта-->

    </li><!-- Конец пункта-->
    <li><a href="https://forum.mhut.org/" class="drop">FORUM</a><!-- Начало пункта-->
        <div class="dropdown_1column">
            <div class="col_1">
                <a href="https://forum.mhut.org/viewtopic.php?p=9000">Sitemap</a>
                <a href="https://forum.mhut.org/viewtopic.php?p=6423">Error report</a>
            </div>

        </div><!-- Конец контейнера-->
    </li><!-- Конец пункта Главная -->
    <li><a href="#" class="drop">DOWNLOAD</a><!-- Начало пункта-->
        <div class="dropdown_4columns"><!-- Начало контейнера-->
            <div class="col_2">
                <h3>Mirrors</h3>
                <a href="http://gen.lib.rus.ec/">Gen.lib.rus.ec</a>
                <a href="http://libgen.lc/">Libgen.lc</a>
                <a href="http://libgen.pw/">Libgen.pw</a>
                <a href="http://b-ok.org/">Z-Library</a>
                <a href="http://bookfi.net/">BookFI.net</a>
            </div>
            <div class="col_1">
                <h3>P2P</h3>
                <a href="/repository_torrent/">Torrents</a>
                <a href="/repository_nzb/">Usenet (*.nzb)</a>
                <h3>Database Dumps</h3>
                <a href="http://gen.lib.rus.ec/dbdumps/">gen.lib.rus.ec</a>
            </div>
            <div class="col_1">
                <h3>Other</h3>
                <a href="https://wiki.mhut.org/software:libgen_desktop">Libgen Desktop application</a>
            </div>
        </div><!-- Конец контейнера-->
    </li><!-- Конец пункта Главная -->
    <li><a href="/librarian/" class="drop">UPLOAD</a><!-- Начало пункта-->
        <div class="dropdown_2columns"><!-- Начало контейнера-->
            <div class="col_2">
                <ul>
                    <a href="/librarian/"><h3>Upload non-fiction content</h3></a>
                    <a href="/foreignfiction/librarian/"><h3>Upload fiction content</h3></a>
                    <a href="/scimag/librarian/"><h3>Upload scientific article</h3></a>

                    (Login:password look at the forum sitemap)
                </ul>
            </div>
        </div><!-- Конец контейнера-->
    </li><!-- Конец пункта-->
    <li><a href="/search.php?mode=last" class="drop">LAST</a><!-- Начало пункта -->
        <div class="dropdown_1column"><!-- Начало контейнера-->
            <div class="col_1">
                <ul>
                    <a href="/search.php?mode=last"><h3>Last added</h3></a>
                    <a href="/search.php?mode=modified">Last modified</a>
                    <a href="/rss/index.php">RSS</a>
                    <a href="https://forum.mhut.org/viewtopic.php?f=17&t=6874">API</a>
                </ul>
            </div>
        </div><!-- Конец контейнера-->
    </li><!-- Конец пункта Главная -->
    <li><a href="#" class="drop">OTHERS</a><!-- Начало пункта-->
        <div class="dropdown_2columns align_right"><!-- Начало контейнера -->
            <div class="col_2">
                <ul>
                    <a href="/comics/"><h3>Comics</h3></a>
                    <a href="/foreignfiction/"><h3>Fiction</h3></a>
                    <a href="http://magzdb.org/"><h3>Magazines</h3></a>
                    <a href="/standarts/"><h3>Standarts</h3></a>
                    <a href="https://b-ok.cc/fulltext/">Full-text search in LG content</a>
                </ul>
            </div>
        </div><!-- Конец контейнера-->
    </li><!-- Конец пункта -->
    <li><a href="#" class="drop">TOPICS</a><!-- Начало пункта-->
        <div class="dropdown_5columns align_right"><!-- Начало контейнера-->
            <div class="col_1">
                <ul class="greybox">
                    <li><a href="../search.php?req=topicid210&open=0&column=topic" class="drop">Technology</a>
                        <ul class="submenu_rightalign">
                            <div class="dropdown_6columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid212&open=0&column=topic">Aerospace Equipment</a></li>
                                    <li><a href="../search.php?req=topicid211&open=0&column=topic">Automation</a></li>
                                    <li><a href="../search.php?req=topicid235&open=0&column=topic">Communication: Telecommunications</a></li>
                                    <li><a href="../search.php?req=topicid234&open=0&column=topic">Communication</a></li>
                                    <li><a href="../search.php?req=topicid236&open=0&column=topic">Construction</a></li>
                                    <li><a href="../search.php?req=topicid241&open=0&column=topic">Construction: Cement Industry</a></li>
                                    <li><a href="../search.php?req=topicid240&open=0&column=topic">Construction: Renovation and interior design: Saunas</a></li>
                                    <li><a href="../search.php?req=topicid239&open=0&column=topic">Construction: Renovation and interior design</a></li>



                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid238&open=0&column=topic">Construction: Ventilation and Air Conditioning</a></li>
                                    <li><a href="../search.php?req=topicid261&open=0&column=topic">Electronics: Electronics</a></li>
                                    <li><a href="../search.php?req=topicid252&open=0&column=topic">Electronics: Fiber Optics</a></li>
                                    <li><a href="../search.php?req=topicid251&open=0&column=topic">Electronics: Hardware</a></li>
                                    <li><a href="../search.php?req=topicid253&open=0&column=topic">Electronics: Home Electronics</a></li>
                                    <li><a href="../search.php?req=topicid254&open=0&column=topic">Electronics: Microprocessor Technology</a></li>
                                    <li><a href="../search.php?req=topicid256&open=0&column=topic">Electronics: Radio</a></li>


                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid257&open=0&column=topic">Electronics: Robotics</a></li>
                                    <li><a href="../search.php?req=topicid255&open=0&column=topic">Electronics: Signal Processing</a></li>
                                    <li><a href="../search.php?req=topicid260&open=0&column=topic">Electronics: Telecommunications</a></li>
                                    <li><a href="../search.php?req=topicid259&open=0&column=topic">Electronics: TV. Video</a></li>
                                    <li><a href="../search.php?req=topicid258&open=0&column=topic">Electronics: VLSI</a></li>
                                    <li><a href="../search.php?req=topicid250&open=0&column=topic">Electronics</a></li>
                                    <li><a href="../search.php?req=topicid263&open=0&column=topic">Energy: Renewable Energy</a></li>
                                    <li><a href="../search.php?req=topicid262&open=0&column=topic">Energy</a></li>
                                    <li><a href="../search.php?req=topicid229&open=0&column=topic">Food Manufacturing</a></li>


                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid243&open=0&column=topic">Fuel Technology</a></li>
                                    <li><a href="../search.php?req=topicid242&open=0&column=topic">Heat</a></li>
                                    <li><a href="../search.php?req=topicid232&open=0&column=topic">industrial equipment and technology</a></li>
                                    <li><a href="../search.php?req=topicid231&open=0&column=topic">Industry: Metallurgy</a></li>
                                    <li><a href="../search.php?req=topicid230&open=0&column=topic">Instrument</a></li>
                                    <li><a href="../search.php?req=topicid218&open=0&column=topic">Light Industry</a></li>
                                    <li><a href="../search.php?req=topicid219&open=0&column=topic">Materials</a></li>
                                    <li><a href="../search.php?req=topicid220&open=0&column=topic">Mechanical Engineering</a></li>
                                    <li><a href="../search.php?req=topicid221&open=0&column=topic">Metallurgy</a></li>
                                    <li><a href="../search.php?req=topicid222&open=0&column=topic">Metrology</a></li>



                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid215&open=0&column=topic">Military equipment: Weapon</a></li>
                                    <li><a href="../search.php?req=topicid214&open=0&column=topic">Military equipment</a></li>
                                    <li><a href="../search.php?req=topicid233&open=0&column=topic">Missiles</a></li>
                                    <li><a href="../search.php?req=topicid224&open=0&column=topic">Nanotechnology</a></li>
                                    <li><a href="../search.php?req=topicid226&open=0&column=topic">Oil and Gas Technologies: Pipelines</a></li>
                                    <li><a href="../search.php?req=topicid225&open=0&column=topic">Oil and Gas Technologies</a></li>
                                    <li><a href="../search.php?req=topicid228&open=0&column=topic">Patent Business. Ingenuity. Innovation</a></li>
                                    <li><a href="../search.php?req=topicid216&open=0&column=topic">Publishing</a></li>
                                    <li><a href="../search.php?req=topicid249&open=0&column=topic">Refrigeration</a></li>
                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid227&open=0&column=topic">Regulatory Literature</a></li>
                                    <li><a href="../search.php?req=topicid223&open=0&column=topic">Safety and Security</a></li>
                                    <li><a href="../search.php?req=topicid217&open=0&column=topic">Space Science</a></li>
                                    <li><a href="../search.php?req=topicid244&open=0&column=topic">Transport</a></li>
                                    <li><a href="../search.php?req=topicid245&open=0&column=topic">Transportation: Aviation</a></li>
                                    <li><a href="../search.php?req=topicid246&open=0&column=topic">Transportation: Cars, motorcycles</a></li>
                                    <li><a href="../search.php?req=topicid247&open=0&column=topic">Transportation: Rail</a></li>
                                    <li><a href="../search.php?req=topicid248&open=0&column=topic">Transportation: Ships</a></li>
                                    <li><a href="../search.php?req=topicid213&open=0&column=topic">Water Treatment</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid57&open=0&column=topic" class="drop">Art</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid60&open=0&column=topic">Cinema</a></li>
                                    <li><a href="../search.php?req=topicid58&open=0&column=topic">Design: Architecture</a></li>
                                    <li><a href="../search.php?req=topicid59&open=0&column=topic">Graphic Arts</a></li>
                                    <li><a href="../search.php?req=topicid61&open=0&column=topic">Music</a></li>
                                    <li><a href="../search.php?req=topicid62&open=0&column=topic">Music: Guitar</a></li>
                                    <li><a href="../search.php?req=topicid63&open=0&column=topic">Photo</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid12&open=0&column=topic" class="drop">Biology</a>
                        <ul>
                            <div class="dropdown_3columns"><!-- Начало контейнера-->
                                <div class="col_1">



                                    <li><a href="../search.php?req=topicid14&open=0&column=topic">Anthropology</a></li>
                                    <li><a href="../search.php?req=topicid15&open=0&column=topic">Anthropology: Evolution</a></li>
                                    <li><a href="../search.php?req=topicid16&open=0&column=topic">Biostatistics</a></li>
                                    <li><a href="../search.php?req=topicid17&open=0&column=topic">Biotechnology</a></li>
                                    <li><a href="../search.php?req=topicid18&open=0&column=topic">Biophysics</a></li>
                                    <li><a href="../search.php?req=topicid19&open=0&column=topic">Biochemistry</a></li>

                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid20&open=0&column=topic">Biochemistry: enologist</a></li>
                                    <li><a href="../search.php?req=topicid31&open=0&column=topic">Ecology</a></li>
                                    <li><a href="../search.php?req=topicid13&open=0&column=topic">Estestvoznananie</a></li>
                                    <li><a href="../search.php?req=topicid22&open=0&column=topic">Genetics</a></li>
                                    <li><a href="../search.php?req=topicid26&open=0&column=topic">Microbiology</a></li>
                                    <li><a href="../search.php?req=topicid27&open=0&column=topic">Molecular</a></li>


                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid28&open=0&column=topic">Molecular: Bioinformatics</a></li>
                                    <li><a href="../search.php?req=topicid30&open=0&column=topic">Plants: Agriculture and Forestry</a></li>
                                    <li><a href="../search.php?req=topicid21&open=0&column=topic">Virology</a></li>
                                    <li><a href="../search.php?req=topicid23&open=0&column=topic">Zoology</a></li>
                                    <li><a href="../search.php?req=topicid24&open=0&column=topic">Zoology:Paleontology</a></li>
                                    <li><a href="../search.php?req=topicid25&open=0&column=topic">Zoology: Fish</a></li>
                                </div>


                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid1&open=0&column=topic" class="drop">Business</a>
                        <ul>
                            <div class="dropdown_2columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid2&open=0&column=topic">Accounting</a></li>
                                    <li><a href="../search.php?req=topicid11&open=0&column=topic">E-Commerce</a></li>
                                    <li><a href="../search.php?req=topicid3&open=0&column=topic">Logistics</a></li>
                                    <li><a href="../search.php?req=topicid6&open=0&column=topic">Management</a></li>
                                    <li><a href="../search.php?req=topicid4&open=0&column=topic">Marketing</a></li>
                                    <li><a href="../search.php?req=topicid5&open=0&column=topic">Marketing: Advertising</a></li>
                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid7&open=0&column=topic">Management: Project Management</a></li>
                                    <li><a href="../search.php?req=topicid8&open=0&column=topic">MLM</a></li>
                                    <li><a href="../search.php?req=topicid9&open=0&column=topic">Responsibility and Business Ethics</a></li>
                                    <li><a href="../search.php?req=topicid10&open=0&column=topic">Trading</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid296&open=0&column=topic" class="drop">Chemistry</a>
                        <ul>
                            <div class="dropdown_2columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid297&open=0&column=topic">Analytical Chemistry</a></li>
                                    <li><a href="../search.php?req=topicid304&open=0&column=topic">Chemical</a></li>
                                    <li><a href="../search.php?req=topicid299&open=0&column=topic">Inorganic Chemistry</a></li>
                                    <li><a href="../search.php?req=topicid298&open=0&column=topic">Materials</a></li>



                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid300&open=0&column=topic">Organic Chemistry</a></li>
                                    <li><a href="../search.php?req=topicid301&open=0&column=topic">Pyrotechnics and explosives</a></li>
                                    <li><a href="../search.php?req=topicid302&open=0&column=topic">Pharmacology</a></li>
                                    <li><a href="../search.php?req=topicid303&open=0&column=topic">Physical Chemistry</a></li>

                                </div>

                            </div>
                        </ul>
                    </li>

                </ul>
            </div>
            <div class="col_1">
                <ul class="greybox">
                    <li><a href="../search.php?req=topicid69&open=0&column=topic" class="drop">Computers</a>
                        <ul>
                            <div class="dropdown_4columns"><!-- Начало контейнера-->
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid71&open=0&column=topic">Algorithms and Data Structures</a></li>
                                    <li><a href="../search.php?req=topicid72&open=0&column=topic">Algorithms and Data Structures: Cryptography</a></li>
                                    <li><a href="../search.php?req=topicid73&open=0&column=topic">Algorithms and Data Structures: Image Processing</a></li>
                                    <li><a href="../search.php?req=topicid74&open=0&column=topic">Algorithms and Data Structures: Pattern Recognition</a></li>
                                    <li><a href="../search.php?req=topicid75&open=0&column=topic">Algorithms and Data Structures: Digital watermarks</a></li>
                                    <li><a href="../search.php?req=topicid80&open=0&column=topic">Cybernetics</a></li>
                                    <li><a href="../search.php?req=topicid81&open=0&column=topic">Cybernetics: ArtificialIntelligence</a></li>
                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid82&open=0&column=topic">Cryptography</a></li>
                                    <li><a href="../search.php?req=topicid76&open=0&column=topic">Databases</a></li>
                                    <li><a href="../search.php?req=topicid78&open=0&column=topic">Information Systems</a></li>
                                    <li><a href="../search.php?req=topicid79&open=0&column=topic">Information Systems: EC businesses</a></li>
                                    <li><a href="../search.php?req=topicid83&open=0&column=topic">Lectures, monographs</a></li>
                                    <li><a href="../search.php?req=topicid84&open=0&column=topic">Media</a></li>
                                    <li><a href="../search.php?req=topicid99&open=0&column=topic">Networking</a></li>
                                    <li><a href="../search.php?req=topicid100&open=0&column=topic">Networking: Internet</a></li>
                                    <li><a href="../search.php?req=topicid85&open=0&column=topic">Operating Systems</a></li>
                                </div>
                                <div class="col_1">



                                    <li><a href="../search.php?req=topicid86&open=0&column=topic">Organization and Data Processing</a></li>
                                    <li><a href="../search.php?req=topicid87&open=0&column=topic">Programming</a></li>
                                    <li><a href="../search.php?req=topicid88&open=0&column=topic">Programming: Libraries API</a></li>
                                    <li><a href="../search.php?req=topicid89&open=0&column=topic">Programming: Games</a></li>
                                    <li><a href="../search.php?req=topicid90&open=0&column=topic">Programming: Compilers</a></li>
                                    <li><a href="../search.php?req=topicid91&open=0&column=topic">Programming: Modeling languages</a></li>
                                    <li><a href="../search.php?req=topicid92&open=0&column=topic">Programming: Programming Languages</a></li>
                                    <li><a href="../search.php?req=topicid93&open=0&column=topic">Programs: TeX, LaTeX</a></li>

                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid77&open=0&column=topic">Security</a></li>
                                    <li><a href="../search.php?req=topicid94&open=0&column=topic">Software: Office software</a></li>
                                    <li><a href="../search.php?req=topicid95&open=0&column=topic">Software: Adobe Products</a></li>
                                    <li><a href="../search.php?req=topicid96&open=0&column=topic">Software: Macromedia Products</a></li>
                                    <li><a href="../search.php?req=topicid97&open=0&column=topic">Software: CAD</a></li>
                                    <li><a href="../search.php?req=topicid98&open=0&column=topic">Software: Systems: scientific computing</a></li>
                                    <li><a href="../search.php?req=topicid101&open=0&column=topic">System Administration</a></li>
                                    <li><a href="../search.php?req=topicid70&open=0&column=topic">Web-design</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid32&open=0&column=topic" class="drop">Geography</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid33&open=0&column=topic">Geodesy. Cartography</a></li>
                                    <li><a href="../search.php?req=topicid34&open=0&column=topic">Local History</a></li>
                                    <li><a href="../search.php?req=topicid35&open=0&column=topic">Local history: Tourism</a></li>
                                    <li><a href="../search.php?req=topicid36&open=0&column=topic">Meteorology, Climatology</a></li>
                                    <li><a href="../search.php?req=topicid37&open=0&column=topic">Russia</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>

                    <li><a href="../search.php?req=topicid38&open=0&column=topic" class="drop">Geology</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid39&open=0&column=topic">Gidrogeology</a></li>
                                    <li><a href="../search.php?req=topicid40&open=0&column=topic">Mining</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid305&open=0&column=topic" class="drop">Economy</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid310&open=0&column=topic">Econometrics</a></li>
                                    <li><a href="../search.php?req=topicid306&open=0&column=topic">Investing</a></li>
                                    <li><a href="../search.php?req=topicid309&open=0&column=topic">Markets</a></li>
                                    <li><a href="../search.php?req=topicid307&open=0&column=topic">Mathematical Economics</a></li>
                                    <li><a href="../search.php?req=topicid308&open=0&column=topic">Popular</a></li>



                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid183&open=0&column=topic" class="drop">Education</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid187&open=0&column=topic">Elementary</a></li>
                                    <li><a href="../search.php?req=topicid185&open=0&column=topic">International Conferences and Symposiums</a></li>
                                    <li><a href="../search.php?req=topicid186&open=0&column=topic">Self-help books</a></li>
                                    <li><a href="../search.php?req=topicid184&open=0&column=topic">Theses abstracts</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>


                </ul>
            </div>
            <div class="col_1">
                <ul class="greybox">
                    <li><a href="../search.php?req=topicid324&open=0&column=topic" class="drop">Jurisprudence</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid311&open=0&column=topic">Criminology, Forensic Science</a></li>
                                    <li><a href="../search.php?req=topicid312&open=0&column=topic">Criminology: Court. examination</a></li>
                                    <li><a href="../search.php?req=topicid313&open=0&column=topic">Law</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid41&open=0&column=topic" class="drop">Housekeeping, leisure</a>
                        <ul>
                            <div class="dropdown_2columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid42&open=0&column=topic">Aquaria</a></li>
                                    <li><a href="../search.php?req=topicid43&open=0&column=topic">Astrology</a></li>
                                    <li><a href="../search.php?req=topicid48&open=0&column=topic">Beauty, image</a></li>
                                    <li><a href="../search.php?req=topicid52&open=0&column=topic">Benefits Homebrew</a></li>
                                    <li><a href="../search.php?req=topicid47&open=0&column=topic">Collecting</a></li>
                                    <li><a href="../search.php?req=topicid49&open=0&column=topic">Cooking</a></li>
                                    <li><a href="../search.php?req=topicid50&open=0&column=topic">Fashion, Jewelry</a></li>
                                    <li><a href="../search.php?req=topicid45&open=0&column=topic">Games: Board Games</a></li>






                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid46&open=0&column=topic">Games: Chess</a></li>
                                    <li><a href="../search.php?req=topicid56&open=0&column=topic">Garden, garden</a></li>
                                    <li><a href="../search.php?req=topicid54&open=0&column=topic">Handicraft</a></li>
                                    <li><a href="../search.php?req=topicid55&open=0&column=topic">Handicraft: Cutting and Sewing</a></li>
                                    <li><a href="../search.php?req=topicid51&open=0&column=topic">Hunting and Game Management</a></li>
                                    <li><a href="../search.php?req=topicid44&open=0&column=topic">Pet</a></li>
                                    <li><a href="../search.php?req=topicid53&open=0&column=topic">Professions and Trades</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid64&open=0&column=topic" class="drop">History</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid65&open=0&column=topic">American Studies</a></li>
                                    <li><a href="../search.php?req=topicid66&open=0&column=topic">Archaeology</a></li>
                                    <li><a href="../search.php?req=topicid67&open=0&column=topic">Military History</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>

                    <li><a href="../search.php?req=topicid314&open=0&column=topic" class="drop">Linguistics</a>
                        <ul>
                            <div class="dropdown_2columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid318&open=0&column=topic">Comparative Studies</a></li>
                                    <li><a href="../search.php?req=topicid322&open=0&column=topic">Dictionaries</a></li>
                                    <li><a href="../search.php?req=topicid315&open=0&column=topic">Foreign</a></li>
                                    <li><a href="../search.php?req=topicid316&open=0&column=topic">Foreign: English</a></li>
                                    <li><a href="../search.php?req=topicid317&open=0&column=topic">Foreign: French</a></li>



                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid319&open=0&column=topic">Linguistics</a></li>
                                    <li><a href="../search.php?req=topicid320&open=0&column=topic">Rhetoric</a></li>
                                    <li><a href="../search.php?req=topicid321&open=0&column=topic">Russian Language</a></li>
                                    <li><a href="../search.php?req=topicid323&open=0&column=topic">Stylistics</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>


                </ul>
            </div>
            <div class="col_1">
                <ul class="greybox">
                    <li><a href="../search.php?req=topicid102&open=0&column=topic" class="drop">Literature</a>
                        <ul>
                            <div class="dropdown_2columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid106&open=0&column=topic">Children</a></li>
                                    <li><a href="../search.php?req=topicid107&open=0&column=topic">Comics</a></li>
                                    <li><a href="../search.php?req=topicid105&open=0&column=topic">Detective</a></li>
                                    <li><a href="../search.php?req=topicid112&open=0&column=topic">Fantasy</a></li>
                                    <li><a href="../search.php?req=topicid103&open=0&column=topic">Fiction</a></li>

                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid111&open=0&column=topic">Folklore</a></li>
                                    <li><a href="../search.php?req=topicid104&open=0&column=topic">Library</a></li>										<li><a href="../search.php?req=topicid108&open=0&column=topic">Literary</a></li>
                                    <li><a href="../search.php?req=topicid109&open=0&column=topic">Poetry</a></li>
                                    <li><a href="../search.php?req=topicid110&open=0&column=topic">Prose</a></li>


                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid113&open=0&column=topic" class="drop">Mathematics</a>
                        <ul>
                            <div class="dropdown_4columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid114&open=0&column=topic">Algebra</a></li>
                                    <li><a href="../search.php?req=topicid115&open=0&column=topic">Algebra: Linear Algebra</a></li>
                                    <li><a href="../search.php?req=topicid116&open=0&column=topic">Algorithms and Data Structures</a></li>
                                    <li><a href="../search.php?req=topicid117&open=0&column=topic">Analysis</a></li>
                                    <li><a href="../search.php?req=topicid137&open=0&column=topic">Applied Mathematics</a></li>
                                    <li><a href="../search.php?req=topicid139&open=0&column=topic">Automatic Control Theory</a></li>
                                    <li><a href="../search.php?req=topicid126&open=0&column=topic">Combinatorics</a></li>
                                    <li><a href="../search.php?req=topicid120&open=0&column=topic">Computational Mathematics</a></li>

                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid128&open=0&column=topic">Computer Algebra</a></li>
                                    <li><a href="../search.php?req=topicid133&open=0&column=topic">Continued fractions</a></li>
                                    <li><a href="../search.php?req=topicid125&open=0&column=topic">Differential Equations</a></li>
                                    <li><a href="../search.php?req=topicid124&open=0&column=topic">Discrete Mathematics</a></li>
                                    <li><a href="../search.php?req=topicid123&open=0&column=topic">Dynamical Systems</a></li>
                                    <li><a href="../search.php?req=topicid146&open=0&column=topic">Elementary</a></li>
                                    <li><a href="../search.php?req=topicid144&open=0&column=topic">Functional Analysis</a></li>
                                    <li><a href="../search.php?req=topicid134&open=0&column=topic">Fuzzy Logic and Applications</a></li>
                                    <li><a href="../search.php?req=topicid141&open=0&column=topic">Game Theory</a></li>

                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid121&open=0&column=topic">Geometry and Topology</a></li>
                                    <li><a href="../search.php?req=topicid140&open=0&column=topic">Graph Theory</a></li>
                                    <li><a href="../search.php?req=topicid129&open=0&column=topic">Lectures</a></li>
                                    <li><a href="../search.php?req=topicid130&open=0&column=topic">Logic</a></li>
                                    <li><a href="../search.php?req=topicid132&open=0&column=topic">Mathematical Physics</a></li>
                                    <li><a href="../search.php?req=topicid131&open=0&column=topic">Mathematical Statistics</a></li>
                                    <li><a href="../search.php?req=topicid143&open=0&column=topic">Number Theory</a></li>
                                    <li><a href="../search.php?req=topicid145&open=0&column=topic">Numerical Analysis</a></li>
                                    <li><a href="../search.php?req=topicid142&open=0&column=topic">Operator Theory</a></li>
                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid135&open=0&column=topic">Optimal control</a></li>
                                    <li><a href="../search.php?req=topicid136&open=0&column=topic">Optimization. Operations Research.</a></li>
                                    <li><a href="../search.php?req=topicid119&open=0&column=topic">Probability</a></li>
                                    <li><a href="../search.php?req=topicid122&open=0&column=topic">Puzzle</a></li>
                                    <li><a href="../search.php?req=topicid138&open=0&column=topic">Symmetry and group</a></li>
                                    <li><a href="../search.php?req=topicid127&open=0&column=topic">The complex variable</a></li>
                                    <li><a href="../search.php?req=topicid118&open=0&column=topic">Wavelets and signal processing</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>

                    <li><a href="../search.php?req=topicid147&open=0&column=topic" class="drop">Medicine</a>
                        <ul>
                            <div class="dropdown_4columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid148&open=0&column=topic">Anatomy and physiology</a></li>
                                    <li><a href="../search.php?req=topicid149&open=0&column=topic">Anesthesiology and Intensive Care</a></li>
                                    <li><a href="../search.php?req=topicid159&open=0&column=topic">Cardiology</a></li>
                                    <li><a href="../search.php?req=topicid160&open=0&column=topic">Chinese Medicine</a></li>
                                    <li><a href="../search.php?req=topicid161&open=0&column=topic">Clinical Medicine</a></li>
                                    <li><a href="../search.php?req=topicid170&open=0&column=topic">Dentistry, Orthodontics</a></li>



                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid155&open=0&column=topic">Diabetes</a></li>
                                    <li><a href="../search.php?req=topicid151&open=0&column=topic">Diseases: Internal Medicine</a></li>
                                    <li><a href="../search.php?req=topicid150&open=0&column=topic">Diseases</a></li>
                                    <li><a href="../search.php?req=topicid176&open=0&column=topic">Endocrinology</a></li>
                                    <li><a href="../search.php?req=topicid167&open=0&column=topic">ENT</a></li>
                                    <li><a href="../search.php?req=topicid177&open=0&column=topic">Epidemiology</a></li>
                                    <li><a href="../search.php?req=topicid174&open=0&column=topic">Feng Shui</a></li>
                                    <li><a href="../search.php?req=topicid152&open=0&column=topic">Histology</a></li>

                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid153&open=0&column=topic">Homeopathy</a></li>
                                    <li><a href="../search.php?req=topicid156&open=0&column=topic">Immunology</a></li>
                                    <li><a href="../search.php?req=topicid157&open=0&column=topic">Infectious diseases</a></li>
                                    <li><a href="../search.php?req=topicid162&open=0&column=topic">Molecular Medicine</a></li>
                                    <li><a href="../search.php?req=topicid163&open=0&column=topic">Natural Medicine</a></li>
                                    <li><a href="../search.php?req=topicid165&open=0&column=topic">Neurology</a></li>
                                    <li><a href="../search.php?req=topicid166&open=0&column=topic">Oncology</a></li>
                                    <li><a href="../search.php?req=topicid168&open=0&column=topic">Ophthalmology</a></li>
                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid169&open=0&column=topic">Pediatrics</a></li>
                                    <li><a href="../search.php?req=topicid173&open=0&column=topic">Pharmacology</a></li>
                                    <li><a href="../search.php?req=topicid164&open=0&column=topic">Popular scientific literature</a></li>
                                    <li><a href="../search.php?req=topicid175&open=0&column=topic">Surgery, Orthopedics</a></li>
                                    <li><a href="../search.php?req=topicid172&open=0&column=topic">Therapy</a></li>
                                    <li><a href="../search.php?req=topicid171&open=0&column=topic">Trial</a></li>
                                    <li><a href="../search.php?req=topicid158&open=0&column=topic">Yoga</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid189&open=0&column=topic" class="drop">Other Social Sciences</a>
                        <ul>
                            <div class="dropdown_2columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid191&open=0&column=topic">Cultural</a></li>
                                    <li><a href="../search.php?req=topicid197&open=0&column=topic">Ethnography</a></li>
                                    <li><a href="../search.php?req=topicid190&open=0&column=topic">Journalism, Media</a></li>
                                    <li><a href="../search.php?req=topicid192&open=0&column=topic">Politics</a></li>
                                    <li><a href="../search.php?req=topicid193&open=0&column=topic">Politics: International Relations</a></li>
                                </div>
                                <div class="col_1">


                                    <li><a href="../search.php?req=topicid195&open=0&column=topic">Philosophy</a></li>
                                    <li><a href="../search.php?req=topicid196&open=0&column=topic">Philosophy: Critical Thinking</a></li>
                                    <li><a href="../search.php?req=topicid194&open=0&column=topic">Sociology</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>

                    <li><a href="../search.php?req=topicid264&open=0&column=topic" class="drop">Physics</a>
                        <ul>
                            <div class="dropdown_4columns"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid266&open=0&column=topic">Astronomy: Astrophysics</a></li>
                                    <li><a href="../search.php?req=topicid265&open=0&column=topic">Astronomy</a></li>
                                    <li><a href="../search.php?req=topicid270&open=0&column=topic">Crystal Physics</a></li>
                                    <li><a href="../search.php?req=topicid287&open=0&column=topic">Electricity and Magnetism</a></li>
                                    <li><a href="../search.php?req=topicid288&open=0&column=topic">Electrodynamics</a></li>
                                    <li><a href="../search.php?req=topicid278&open=0&column=topic">General courses</a></li>
                                    <li><a href="../search.php?req=topicid267&open=0&column=topic">Geophysics</a></li>
                                    <li><a href="../search.php?req=topicid271&open=0&column=topic">Mechanics</a></li>

                                </div>
                                <div class="col_1">

                                    <li><a href="../search.php?req=topicid274&open=0&column=topic">Mechanics: Fluid Mechanics</a></li>
                                    <li><a href="../search.php?req=topicid273&open=0&column=topic">Mechanics: Mechanics of deformable bodies</a></li>
                                    <li><a href="../search.php?req=topicid275&open=0&column=topic">Mechanics: Nonlinear dynamics and chaos</a></li>
                                    <li><a href="../search.php?req=topicid272&open=0&column=topic">Mechanics: Oscillations and Waves</a></li>




                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid276&open=0&column=topic">Mechanics: Strength of Materials</a></li>
                                    <li><a href="../search.php?req=topicid277&open=0&column=topic">Mechanics: Theory of Elasticity</a></li>
                                    <li><a href="../search.php?req=topicid279&open=0&column=topic">Optics</a></li>
                                    <li><a href="../search.php?req=topicid284&open=0&column=topic">Physics of lasers</a></li>
                                    <li><a href="../search.php?req=topicid283&open=0&column=topic">Physics of the Atmosphere</a></li>
                                    <li><a href="../search.php?req=topicid285&open=0&column=topic">Plasma Physics</a></li>


                                </div>
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid268&open=0&column=topic">Quantum Mechanics</a></li>
                                    <li><a href="../search.php?req=topicid269&open=0&column=topic">Quantum Physics</a></li>
                                    <li><a href="../search.php?req=topicid286&open=0&column=topic">Solid State Physics</a></li>
                                    <li><a href="../search.php?req=topicid280&open=0&column=topic">Spectroscopy</a></li>
                                    <li><a href="../search.php?req=topicid281&open=0&column=topic">Theory of Relativity and Gravitation</a></li>
                                    <li><a href="../search.php?req=topicid282&open=0&column=topic">Thermodynamics and Statistical Mechanics</a></li>
                                </div>
                            </div>
                        </ul>
                    </li>

                </ul>
            </div>
            <div class="col_1">
                <ul class="greybox">
                    <li><a href="../search.php?req=topicid289&open=0&column=topic" class="drop">Physical Educ. and Sport</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid290&open=0&column=topic">Bodybuilding</a></li>
                                    <li><a href="../search.php?req=topicid292&open=0&column=topic">Bike</a></li>
                                    <li><a href="../search.php?req=topicid295&open=0&column=topic">Fencing</a></li>
                                    <li><a href="../search.php?req=topicid291&open=0&column=topic">Martial Arts</a></li>
                                    <li><a href="../search.php?req=topicid294&open=0&column=topic">Sport fishing</a></li>
                                    <li><a href="../search.php?req=topicid293&open=0&column=topic">Survival</a></li>



                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid198&open=0&column=topic" class="drop">Psychology</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid200&open=0&column=topic">The art of communication</a></li>
                                    <li><a href="../search.php?req=topicid204&open=0&column=topic">Creative Thinking</a></li>
                                    <li><a href="../search.php?req=topicid199&open=0&column=topic">Hypnosis</a></li>
                                    <li><a href="../search.php?req=topicid201&open=0&column=topic">Love, erotic</a></li>
                                    <li><a href="../search.php?req=topicid202&open=0&column=topic">Neuro-Linguistic Programming</a></li>
                                    <li><a href="../search.php?req=topicid203&open=0&column=topic">Pedagogy</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid205&open=0&column=topic" class="drop">Religion</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid206&open=0&column=topic">Buddhism</a></li>
                                    <li><a href="../search.php?req=topicid209&open=0&column=topic">Esoteric, Mystery</a></li>
                                    <li><a href="../search.php?req=topicid207&open=0&column=topic">Kabbalah</a></li>
                                    <li><a href="../search.php?req=topicid208&open=0&column=topic">Orthodoxy</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>
                    <li><a href="../search.php?req=topicid178&open=0&column=topic" class="drop">Science (General)</a>
                        <ul>
                            <div class="dropdown_1column"><!-- Начало контейнера-->
                                <div class="col_1">
                                    <li><a href="../search.php?req=topicid179&open=0&column=topic">International Conferences and Symposiums</a></li>
                                    <li><a href="../search.php?req=topicid180&open=0&column=topic">Science of Science</a></li>
                                    <li><a href="../search.php?req=topicid181&open=0&column=topic">Scientific-popular</a></li>
                                    <li><a href="../search.php?req=topicid182&open=0&column=topic">Scientific and popular: Journalism</a></li>

                                </div>
                            </div>
                        </ul>
                    </li>

                </ul>
            </div>
        </div><!-- Конец контейнера-->
    </li><!-- Конец пункта-->
</ul>
<link rel='stylesheet' type='text/css' href='../paginator3000.css' />
<script type='text/javascript' src='../paginator3000.js'></script>
<style type='text/css'>
    .c { font-family: Georgia, 'Times New Roman', Times, serif; font-size: 11px; color: #000000; LETTER-SPACING: 0px; }
    A { text-decoration: none; }
    td { padding: 1px; }
    table { border-spacing: 1px 1px; }
</style>
<table width=100% border=0><tr><td><form name ='libgen' action='search.php'><br>
                <input autofocus='autofocus' name='req' id='searchform' size=60 maxlength=80 value='test'>
                <input type=submit onclick='this.disabled='disabled'; document.forms.item(0).submit();' value='Search!'><br>
                <font face=Arial color=gray size=1><a href='../batchsearchindex.php'>Batch search for books</a></font><br>
                <label><b>Download type:</b></label>
                <select name='open' size='1'>
                    <option value='0' selected='selected'>Resumed dl with original filename</option>
                    <option value='1'>Resumed dl with translit filename</option>
                    <option value='2'>Resumed dl with md5 filename</option>
                    <option value='3'>Open file in browser</option>
                </select>
                <label><b>Results per page</b></label>
                <select name='res' size='1'>
                    <option value='25' selected='selected'>25</option>
                    <option value='50'>50</option>
                    <option value='100'>100</option>
                </select>
                <br>
                <b>View results:</b>
                <input type=radio name='view' checked value='simple'>
                <label for='simple'>Simple</label>
                <input type=radio name='view'    value='detailed'>
                <label for='detailed'>Detailed</label>
                <b>   Search with mask (word*):</b>
                <input type=radio name='phrase'  checked  value='1'>
                <label for='detailed'>No</label>
                <input type=radio name='phrase'  value='0'>
                <label for='simple'>Yes</label>
                <br>
                <font><b>Search in fields</b></font>
                <input type='radio' name='column' value='def' checked><a href='#' title='Columns: Title,Author(s),Series,Periodical,Publisher,Year,VolumeInfo'>The column set default</a>
                <input type='radio' name='column' value='title'>Title
                <input type='radio' name='column' value='author'>Author(s)
                <input type='radio' name='column' value='series'>Series<br>

                <input type='radio' name='column' value='publisher'>Publisher
                <input type='radio' name='column' value='year'>Year
                <input type='radio' name='column' value='identifier'>ISBN
                <input type='radio' name='column' value='language'><a href='' title='Russian, English, German, French, Spanish, ... etc. (ISO 639)'>Language</a>
                <input type='radio' name='column' value='md5'>MD5
                <input type='radio' name='column' value='tags'>Tags
                <input type='radio' name='column' value='extension'>Extension
            </form></td><td><h1 style="color:#A00000"><a href="/">Library Genesis<sup style="font-size:65%">2M</sup></h1></a><br/><a href="https://wiki.mhut.org/software:libgen_desktop">Introducing Libgen Desktop application!</a><br><a href="http://custodians.online/">Letter of Solidarity</a></td></tr></table><div style="text-align: center; float: left;" class="paginator" id="paginator_example_top"></div>
<script type="text/javascript">
    paginator_example_top = new Paginator(
        "paginator_example_top", // id контейнера, куда ляжет пагинатор
        40, // общее число страниц
        25, // число страниц, видимых одновременно
        1, // номер текущей страницы
        "search.php?&req=test&phrase=1&view=simple&column=def&sort=def&sortmode=ASC&page=" // url страниц
    );
</script>
<table width=100%><tr><td align='left' width=45%><font color=grey size=1>2780 files found , Showing the first  1000  results | showing results from 1 to 25</font></td><td align=center width=10%><font size="3" color="gray"><a href="search.php?&req=test&phrase=1&view=simple&column=def&sort=def&sortmode=ASC&page=2">&nbsp;&nbsp;&#9658;</a></font></td><td align='right' width=45%><font size=2>also search "test"  in   <a href='/foreignfiction/index.php?s=test&f_lang=All&f_columns=0&f_ext=All&f_group=1'>fiction</a></font></td></tr></table><table width=100% cellspacing=1 cellpadding=1 rules=rows class=c align=center><tr valign=top bgcolor=#C0C0C0>
        <td><b><a title='Sort results by ID' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=id&sortmode=DESC'>ID</a></b></td>
        <td><b><a title='Sort results by Author' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=author&sortmode=DESC'>Author(s)</a></b></td>
        <td><b><a title='Sort results by Title' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=title&sortmode=DESC'>Title</a></b></td>
        <td><b><a title='Sort results by Publisher' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=publisher&sortmode=DESC'>Publisher</a></b></td>
        <td><b><a title='Sort results by Year' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=year&sortmode=DESC'>Year</a></b></td>
        <td><b><a title='Sort results by Pages' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=pages&sortmode=DESC'>Pages</a></b></td>
        <td><b><a title='Sort results by Language' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=language&sortmode=DESC'>Language</a></b></td>
        <td><b><a title='Sort results by Size' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=filesize&sortmode=DESC'>Size</a></b></td>
        <td><b><a title='Sort results by Extension' href='search.php?&req=test&phrase=1&view=simple&column=def&sort=extension&sortmode=DESC'>Extension</a></b></td>
        <td colspan=5><b>Mirrors</b></td>
        <td><b>Edit</b></td></tr><tr valign=top bgcolor=#C6DEFF><td>643</td>
        <td><a href='search.php?req=Larry J. Crockett&column[]=author'>Larry J. Crockett</a></td>
        <td width=500><a href="search.php?req=Ablex+Series+in+Artificial+Intelligence&column=series"><font face=Times color=green><i>Ablex Series in Artificial Intelligence</i></font></a><br><a href='book/index.php?md5=2F2DBA2A621B693BB95601C16ED680F8' title='' id=643>The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence<br> <font face=Times color=green><i>9780893919269, 0893919268</i></font></a></td>
        <td>Ablex Publishing Corporation</td>
        <td nowrap>1994</td>
        <td>216</td>
        <td>English</td>
        <td nowrap>517 Kb</td>
        <td nowrap>gz</td>
        <td><a href='http://93.174.95.29/_ads/2F2DBA2A621B693BB95601C16ED680F8' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=2F2DBA2A621B693BB95601C16ED680F8' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/2F2DBA2A621B693BB95601C16ED680F8' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=643' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/2F2DBA2A621B693BB95601C16ED680F8' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/2F2DBA2A621B693BB95601C16ED680F8' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>3167</td>
        <td><a href='search.php?req=M. Shifman&column[]=author'>M. Shifman</a></td>
        <td width=500><a href='book/index.php?md5=06E6135019C8F2F43158ABA9ABDC610E' title='' id=3167>You failed your math test, Comrade Einstein (about Soviet antisemitism)<br> <font face=Times color=green><i>9789812562791, 9812562796</i></font></a></td>
        <td>World Scientific Publishing Company</td>
        <td nowrap>2005</td>
        <td>268</td>
        <td>English</td>
        <td nowrap>3 Mb</td>
        <td nowrap>djvu</td>
        <td><a href='http://93.174.95.29/_ads/06E6135019C8F2F43158ABA9ABDC610E' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=06E6135019C8F2F43158ABA9ABDC610E' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/06E6135019C8F2F43158ABA9ABDC610E' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=3167' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/06E6135019C8F2F43158ABA9ABDC610E' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/06E6135019C8F2F43158ABA9ABDC610E' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>9996</td>
        <td><a href='search.php?req=Martin Gardner&column[]=author'>Martin Gardner</a></td>
        <td width=500><a href="search.php?req=Test+Your+Code+Breaking+Skills&column=series"><font face=Times color=green><i>Test Your Code Breaking Skills</i></font></a><br><a href='book/index.php?md5=4363AD191DB6B625BC6200326A51E5DF' title='' id=9996>Codes, ciphers, and secret writing<br> <font face=Times color=green><i>9780486247618, 0486247619</i></font></a></td>
        <td>Dover Publications</td>
        <td nowrap>1984</td>
        <td>48</td>
        <td>English</td>
        <td nowrap>725 Kb</td>
        <td nowrap>djvu</td>
        <td><a href='http://93.174.95.29/_ads/4363AD191DB6B625BC6200326A51E5DF' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=4363AD191DB6B625BC6200326A51E5DF' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/4363AD191DB6B625BC6200326A51E5DF' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=9996' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/4363AD191DB6B625BC6200326A51E5DF' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/4363AD191DB6B625BC6200326A51E5DF' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>14585</td>
        <td><a href='search.php?req=&column[]=author'></a></td>
        <td width=500><a href='book/index.php?md5=255C181F183AF07A7B8EA95433927DDE' title='' id=14585>The GRE Physics Test practice book</a></td>
        <td>ETS</td>
        <td nowrap>2001</td>
        <td>79</td>
        <td>English</td>
        <td nowrap>562 Kb</td>
        <td nowrap>djvu</td>
        <td><a href='http://93.174.95.29/_ads/255C181F183AF07A7B8EA95433927DDE' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=255C181F183AF07A7B8EA95433927DDE' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/255C181F183AF07A7B8EA95433927DDE' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=14585' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/255C181F183AF07A7B8EA95433927DDE' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/255C181F183AF07A7B8EA95433927DDE' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>18023</td>
        <td><a href='search.php?req=Joseph Molitoris&column[]=author'>Joseph Molitoris</a></td>
        <td width=500><a href='book/index.php?md5=E185937818FE3285AEDDF19B9A85EF9B' title='' id=18023>The GRE physics test preparation<br> <font face=Times color=green><i>0878918485, 9780878918485</i></font></a></td>
        <td>Research &amp; Education Association</td>
        <td nowrap>1991</td>
        <td>406</td>
        <td>English</td>
        <td nowrap>3 Mb</td>
        <td nowrap>djvu</td>
        <td><a href='http://93.174.95.29/_ads/E185937818FE3285AEDDF19B9A85EF9B' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=E185937818FE3285AEDDF19B9A85EF9B' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/E185937818FE3285AEDDF19B9A85EF9B' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=18023' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/E185937818FE3285AEDDF19B9A85EF9B' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/E185937818FE3285AEDDF19B9A85EF9B' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>880351</td>
        <td><a href='search.php?req=Shannon R. Turlington&column[]=author'>Shannon R. Turlington</a></td>
        <td width=500><a href='book/index.php?md5=580000A1CAA698C2EFD8F5439E9A1F26' title='' id=880351>Master The Civil Service Exam: Targeted Test Prep to Jump-Start Your Career <font face=Times color=green><i>[4&nbsp;ed.]</i></font><br> <font face=Times color=green><i>0768927196, 9780768927191</i></font></a></td>
        <td>Peterson's</td>
        <td nowrap>2009</td>
        <td>520</td>
        <td>English</td>
        <td nowrap>7 Mb</td>
        <td nowrap>epub</td>
        <td><a href='http://93.174.95.29/_ads/580000A1CAA698C2EFD8F5439E9A1F26' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=580000A1CAA698C2EFD8F5439E9A1F26' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/580000A1CAA698C2EFD8F5439E9A1F26' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=880351' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/580000A1CAA698C2EFD8F5439E9A1F26' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/580000A1CAA698C2EFD8F5439E9A1F26' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>22619</td>
        <td><a href='search.php?req=Kent Beck&column[]=author'>Kent Beck</a></td>
        <td width=500><a href="search.php?req=Addison-Wesley+Signature+Series&column=series"><font face=Times color=green><i>Addison-Wesley Signature Series</i></font></a><br><a href='book/index.php?md5=1C0D9E39A4EFE28138A88A5A1BD54D5D' title='' id=22619>Test-driven development by example<br> <font face=Times color=green><i>0321146530, 9780321146533</i></font></a></td>
        <td>Addison-Wesley Professional</td>
        <td nowrap>2002</td>
        <td>240</td>
        <td>English</td>
        <td nowrap>382 Kb</td>
        <td nowrap>chm</td>
        <td><a href='http://93.174.95.29/_ads/1C0D9E39A4EFE28138A88A5A1BD54D5D' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=1C0D9E39A4EFE28138A88A5A1BD54D5D' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/1C0D9E39A4EFE28138A88A5A1BD54D5D' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=22619' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/1C0D9E39A4EFE28138A88A5A1BD54D5D' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/1C0D9E39A4EFE28138A88A5A1BD54D5D' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>23658</td>
        <td><a href="search.php?req=Clyde+F.+Coombs&column=author">Clyde F. Coombs</a>,<a href="search.php?req=+Catherine+Ann+Coombs&column=author"> Catherine Ann Coombs</a></td>
        <td width=500><a href='book/index.php?md5=273B7FFDDEC4D0339A7C9F43643627BD' title='' id=23658>Communications Network Test and Measurement Handbook <font face=Times color=green><i>[1&nbsp;ed.]</i></font><br> <font face=Times color=green><i>9780070126176, 0070126178</i></font></a></td>
        <td>McGraw-Hill Professional</td>
        <td nowrap>1997</td>
        <td>786</td>
        <td>English</td>
        <td nowrap>8 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/273B7FFDDEC4D0339A7C9F43643627BD' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=273B7FFDDEC4D0339A7C9F43643627BD' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/273B7FFDDEC4D0339A7C9F43643627BD' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=23658' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/273B7FFDDEC4D0339A7C9F43643627BD' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/273B7FFDDEC4D0339A7C9F43643627BD' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>24876</td>
        <td><a href="search.php?req=Gold+R.&column=author">Gold R.</a>,<a href="search.php?req=+Hammell+T.&column=author"> Hammell T.</a>,<a href="search.php?req=+Snyder+T.&column=author"> Snyder T.</a></td>
        <td width=500><a href='book/index.php?md5=CFD0B1D96ED02AF5347B3D645D444FD6' title='' id=24876>Test-driven development: A J2EE Example</a></td>
        <td></td>
        <td nowrap>2004</td>
        <td>296</td>
        <td>English</td>
        <td nowrap>2 Mb</td>
        <td nowrap>chm</td>
        <td><a href='http://93.174.95.29/_ads/CFD0B1D96ED02AF5347B3D645D444FD6' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=CFD0B1D96ED02AF5347B3D645D444FD6' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/CFD0B1D96ED02AF5347B3D645D444FD6' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=24876' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/CFD0B1D96ED02AF5347B3D645D444FD6' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/CFD0B1D96ED02AF5347B3D645D444FD6' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>24937</td>
        <td><a href="search.php?req=Gosselin+J.&column=author">Gosselin J.</a>,<a href="search.php?req=+Cloutier+L.&column=author"> Cloutier L.</a></td>
        <td width=500><a href='book/index.php?md5=3E8E669B4D875B0988623AC590A1FE6A' title='' id=24937>Oracle Enterprise Manager. Event Test Reference Manual <font face=Times color=green><i>[release 9.0.1&nbsp;ed.]</i></font></a></td>
        <td></td>
        <td nowrap>2001</td>
        <td>514</td>
        <td>English</td>
        <td nowrap>3 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/3E8E669B4D875B0988623AC590A1FE6A' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=3E8E669B4D875B0988623AC590A1FE6A' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/3E8E669B4D875B0988623AC590A1FE6A' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=24937' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/3E8E669B4D875B0988623AC590A1FE6A' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/3E8E669B4D875B0988623AC590A1FE6A' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>25175</td>
        <td><a href="search.php?req=Paul+Hamill&column=author">Paul Hamill</a></td>
        <td width=500><a href='book/index.php?md5=6D71FE24072880B78B318778CCE114D6' title='' id=25175>Unit Test Frameworks<br> <font face=Times color=green><i>9780596006891, 0596006896</i></font></a></td>
        <td>O'Reilly Media</td>
        <td nowrap>2004</td>
        <td>304</td>
        <td>English</td>
        <td nowrap>587 Kb</td>
        <td nowrap>chm</td>
        <td><a href='http://93.174.95.29/_ads/6D71FE24072880B78B318778CCE114D6' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=6D71FE24072880B78B318778CCE114D6' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/6D71FE24072880B78B318778CCE114D6' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=25175' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/6D71FE24072880B78B318778CCE114D6' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/6D71FE24072880B78B318778CCE114D6' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>25247</td>
        <td><a href="search.php?req=Mark+Harrison&column=author">Mark Harrison</a></td>
        <td width=500><a href='book/index.php?md5=4A5EC43857DD992575E2CD2CF28F030C' title='' id=25247>CPE practice test<br> <font face=Times color=green><i>0194329089, 9780194329088</i></font></a></td>
        <td>Oxford University Press</td>
        <td nowrap>2002</td>
        <td>230</td>
        <td>English</td>
        <td nowrap>3 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/4A5EC43857DD992575E2CD2CF28F030C' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=4A5EC43857DD992575E2CD2CF28F030C' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/4A5EC43857DD992575E2CD2CF28F030C' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=25247' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/4A5EC43857DD992575E2CD2CF28F030C' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/4A5EC43857DD992575E2CD2CF28F030C' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>25759</td>
        <td><a href="search.php?req=Vanessa+Jakeman&column=author">Vanessa Jakeman</a>,<a href="search.php?req=+Clare+McDowell&column=author"> Clare McDowell</a></td>
        <td width=500><a href='book/index.php?md5=EBE88112E9F0738C13CFEFA25EB5F0E5' title='' id=25759>Cambridge practice test for IELTS 1 <font face=Times color=green><i>[abridged edition annotated edition]</i></font><br> <font face=Times color=green><i>9780521497664, 0-521-49766-3</i></font></a></td>
        <td>Cambridge University Press</td>
        <td nowrap>1997</td>
        <td>162</td>
        <td>English</td>
        <td nowrap>9 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/EBE88112E9F0738C13CFEFA25EB5F0E5' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=EBE88112E9F0738C13CFEFA25EB5F0E5' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/EBE88112E9F0738C13CFEFA25EB5F0E5' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=25759' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/EBE88112E9F0738C13CFEFA25EB5F0E5' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/EBE88112E9F0738C13CFEFA25EB5F0E5' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>25760</td>
        <td><a href="search.php?req=Jakeman+V.&column=author">Jakeman V.</a>,<a href="search.php?req=+McDowell+C.&column=author"> McDowell C.</a></td>
        <td width=500><a href='book/index.php?md5=E96D64EDDF1A38A1C8D7F986F1B42E38' title='' id=25760>Cambridge practice test for IELTS 3</a></td>
        <td></td>
        <td nowrap>2002</td>
        <td>179</td>
        <td>English</td>
        <td nowrap>17 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/E96D64EDDF1A38A1C8D7F986F1B42E38' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=E96D64EDDF1A38A1C8D7F986F1B42E38' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/E96D64EDDF1A38A1C8D7F986F1B42E38' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=25760' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/E96D64EDDF1A38A1C8D7F986F1B42E38' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/E96D64EDDF1A38A1C8D7F986F1B42E38' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>26423</td>
        <td><a href="search.php?req=Lamport+L.&column=author">Lamport L.</a></td>
        <td width=500><a href='book/index.php?md5=AF01E2453C813CBADF6BB010144AC36B' title='' id=26423>Test File</a></td>
        <td></td>
        <td nowrap>1984</td>
        <td>160</td>
        <td>English</td>
        <td nowrap>717 Kb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/AF01E2453C813CBADF6BB010144AC36B' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=AF01E2453C813CBADF6BB010144AC36B' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/AF01E2453C813CBADF6BB010144AC36B' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=26423' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/AF01E2453C813CBADF6BB010144AC36B' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/AF01E2453C813CBADF6BB010144AC36B' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>27043</td>
        <td><a href="search.php?req=Dr.+James+McCaffrey&column=author">Dr. James McCaffrey</a></td>
        <td width=500><a href="search.php?req=Expert%27s+Voice+in+.Net&column=series"><font face=Times color=green><i>Expert's Voice in .Net</i></font></a><br><a href='book/index.php?md5=48D49EC713291D7EE1487920FCFA7A26' title='' id=27043>.NET Test Automation Recipes: A Problem-Solution Approach <font face=Times color=green><i>[1&nbsp;ed.]</i></font><br> <font face=Times color=green><i>1590596633, 9781590596630, 9781430201632</i></font></a></td>
        <td>Apress</td>
        <td nowrap>2006</td>
        <td>403</td>
        <td>English</td>
        <td nowrap>2 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/48D49EC713291D7EE1487920FCFA7A26' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=48D49EC713291D7EE1487920FCFA7A26' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/48D49EC713291D7EE1487920FCFA7A26' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=27043' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/48D49EC713291D7EE1487920FCFA7A26' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/48D49EC713291D7EE1487920FCFA7A26' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>27606</td>
        <td><a href="search.php?req=James+W.+Newkirk&column=author">James W. Newkirk</a>,<a href="search.php?req=+Alexei+A.+Vorontsov&column=author"> Alexei A. Vorontsov</a></td>
        <td width=500><a href="search.php?req=Microsoft+Professional&column=series"><font face=Times color=green><i>Microsoft Professional</i></font></a><br><a href='book/index.php?md5=C325594657AD93951257E2AC611896D5' title='' id=27606>Test-Driven Development in Microsoft .NET <font face=Times color=green><i>[1&nbsp;ed.]</i></font><br> <font face=Times color=green><i>9780735619487, 0735619484</i></font></a></td>
        <td>Microsoft Press</td>
        <td nowrap>2004</td>
        <td>304</td>
        <td>English</td>
        <td nowrap>2 Mb</td>
        <td nowrap>chm</td>
        <td><a href='http://93.174.95.29/_ads/C325594657AD93951257E2AC611896D5' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=C325594657AD93951257E2AC611896D5' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/C325594657AD93951257E2AC611896D5' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=27606' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/C325594657AD93951257E2AC611896D5' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/C325594657AD93951257E2AC611896D5' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>27882</td>
        <td><a href="search.php?req=Bruce+Rogers&column=author">Bruce Rogers</a></td>
        <td width=500><a href="search.php?req=Complete+Guide+to+the+Toefl+Test&column=series"><font face=Times color=green><i>Complete Guide to the Toefl Test</i></font></a><br><a href='book/index.php?md5=EBBD9ABA948E60BE50DABCD9919E9011' title='' id=27882>TOEFL Secrets <font face=Times color=green><i>[4th Bk&amp;Cdr-IBT/E&nbsp;ed.]</i></font><br> <font face=Times color=green><i>9781413023039, 1413023037</i></font></a></td>
        <td>Heinle ELT</td>
        <td nowrap>2006</td>
        <td>100</td>
        <td>English</td>
        <td nowrap>1 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/EBBD9ABA948E60BE50DABCD9919E9011' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=EBBD9ABA948E60BE50DABCD9919E9011' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/EBBD9ABA948E60BE50DABCD9919E9011' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=27882' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/EBBD9ABA948E60BE50DABCD9919E9011' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/EBBD9ABA948E60BE50DABCD9919E9011' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>28152</td>
        <td><a href="search.php?req=David+Prutchi&column=author">David Prutchi</a>,<a href="search.php?req=+Michael+Norris&column=author"> Michael Norris</a></td>
        <td width=500><a href='book/index.php?md5=92500491552067AB1EFC548FE049B5B2' title='' id=28152>Design and Development of Medical Electronic Instrumentation: A Practical Perspective of the Design, Construction, and Test of Medical Devices <font face=Times color=green><i>[1&nbsp;ed.]</i></font><br> <font face=Times color=green><i>9780471676232, 0471676233</i></font></a></td>
        <td>Wiley-Interscience</td>
        <td nowrap>2004</td>
        <td>479</td>
        <td>English</td>
        <td nowrap>14 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/92500491552067AB1EFC548FE049B5B2' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=92500491552067AB1EFC548FE049B5B2' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/92500491552067AB1EFC548FE049B5B2' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=28152' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/92500491552067AB1EFC548FE049B5B2' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/92500491552067AB1EFC548FE049B5B2' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>28167</td>
        <td><a href="search.php?req=Michael+A.+Pyle&column=author">Michael A. Pyle</a></td>
        <td width=500><a href="search.php?req=Cliffs+Test+Prep&column=series"><font face=Times color=green><i>Cliffs Test Prep</i></font></a><br><a href='book/index.php?md5=6409E129895F7A1C428B3DDF0AD0EE46' title='' id=28167>CliffsTestPrep. TOEFL CBT<br> <font face=Times color=green><i>0764586092, 9780764586095, 9780764517150</i></font></a></td>
        <td>Cliffs Notes</td>
        <td nowrap>2000</td>
        <td>449</td>
        <td>English</td>
        <td nowrap>2 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/6409E129895F7A1C428B3DDF0AD0EE46' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=6409E129895F7A1C428B3DDF0AD0EE46' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/6409E129895F7A1C428B3DDF0AD0EE46' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=28167' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/6409E129895F7A1C428B3DDF0AD0EE46' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/6409E129895F7A1C428B3DDF0AD0EE46' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>28652</td>
        <td><a href="search.php?req=Stephen+Scheiber&column=author">Stephen Scheiber</a></td>
        <td width=500><a href="search.php?req=Test+and+Measurement+Series&column=series"><font face=Times color=green><i>Test and Measurement Series</i></font></a><br><a href='book/index.php?md5=67B73AE2A1E3FEC2A36F1A3542502A13' title='' id=28652>Building a Successful Board-Test Strategy <font face=Times color=green><i>[2nd ed]</i></font><br> <font face=Times color=green><i>9780750672801, 0-7506-7280-3</i></font></a></td>
        <td>Butterworth-Heinemann</td>
        <td nowrap>2001</td>
        <td>350</td>
        <td>English</td>
        <td nowrap>23 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/67B73AE2A1E3FEC2A36F1A3542502A13' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=67B73AE2A1E3FEC2A36F1A3542502A13' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/67B73AE2A1E3FEC2A36F1A3542502A13' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=28652' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/67B73AE2A1E3FEC2A36F1A3542502A13' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/67B73AE2A1E3FEC2A36F1A3542502A13' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>28664</td>
        <td><a href="search.php?req=Mike+Schiffman&column=author">Mike Schiffman</a></td>
        <td width=500><a href='book/index.php?md5=E7BD260E15A578D011F9127B7B963F37' title='' id=28664>Hacker’s challenge: test your incident response skills using 20 scenarios <font face=Times color=green><i>[First edition.]</i></font><br> <font face=Times color=green><i>0072193840, 9780072193848</i></font></a></td>
        <td>McGraw-Hill Osborne Media</td>
        <td nowrap>2001</td>
        <td>384</td>
        <td>English</td>
        <td nowrap>19 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/E7BD260E15A578D011F9127B7B963F37' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=E7BD260E15A578D011F9127B7B963F37' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/E7BD260E15A578D011F9127B7B963F37' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=28664' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/E7BD260E15A578D011F9127B7B963F37' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/E7BD260E15A578D011F9127B7B963F37' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>34730</td>
        <td><a href="search.php?req=&column=author"></a></td>
        <td width=500><a href='book/index.php?md5=FE48C60A5BF64472E5827562D0DB160A' title='' id=34730>Actual Test: Implementing, Managing and Maintaining a MS Windows Server 2003 Network Infrastructure</a></td>
        <td></td>
        <td nowrap>2004</td>
        <td>81</td>
        <td>English</td>
        <td nowrap>9 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/FE48C60A5BF64472E5827562D0DB160A' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=FE48C60A5BF64472E5827562D0DB160A' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/FE48C60A5BF64472E5827562D0DB160A' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=34730' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/FE48C60A5BF64472E5827562D0DB160A' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/FE48C60A5BF64472E5827562D0DB160A' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=><td>35429</td>
        <td><a href="search.php?req=&column=author"></a></td>
        <td width=500><a href='book/index.php?md5=FE10C0A9117E3D288E199E7E352FCCF9' title='' id=35429>Math Review for Practicing to Take the GRE General Test</a></td>
        <td></td>
        <td nowrap>1994</td>
        <td>68</td>
        <td>English</td>
        <td nowrap>7 Mb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/FE10C0A9117E3D288E199E7E352FCCF9' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=FE10C0A9117E3D288E199E7E352FCCF9' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/FE10C0A9117E3D288E199E7E352FCCF9' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=35429' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/FE10C0A9117E3D288E199E7E352FCCF9' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/FE10C0A9117E3D288E199E7E352FCCF9' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    <tr valign=top bgcolor=#C6DEFF><td>36084</td>
        <td><a href="search.php?req=&column=author"></a></td>
        <td width=500><a href='book/index.php?md5=0B6E27281A0193B4A793F4EF4810609F' title='' id=36084>Sun Certified Solaris 8. Administrator Test. Brain Dumps <font face=Times color=green><i>[Part 1]</i></font></a></td>
        <td></td>
        <td nowrap></td>
        <td>28</td>
        <td>English</td>
        <td nowrap>55 Kb</td>
        <td nowrap>pdf</td>
        <td><a href='http://93.174.95.29/_ads/0B6E27281A0193B4A793F4EF4810609F' title='Gen.lib.rus.ec'>[1]</a></td><td><a href='http://libgen.lc/ads.php?md5=0B6E27281A0193B4A793F4EF4810609F' title='Libgen.lc'>[2]</a></td><td><a href='http://b-ok.cc/md5/0B6E27281A0193B4A793F4EF4810609F' title='Z-Library'>[3]</a></td><td><a href='https://libgen.pw/item?id=36084' title='Libgen.pw'>[4]</a></td><td><a href='http://bookfi.net/md5/0B6E27281A0193B4A793F4EF4810609F' title='BookFI.net'>[5]</a></td>
        <td><a href='https://library.bz/main/edit/0B6E27281A0193B4A793F4EF4810609F' title='Libgen Librarian'>[edit]</a></td>
    </tr>

    </tr></table>
<div style="text-align: center;" class="paginator" id="paginator_example_bottom"></div>
<script type="text/javascript">
    paginator_example_bottom = new Paginator(
        "paginator_example_bottom", // id контейнера, куда ляжет пагинатор
        40, // общее число страниц
        25, // число страниц, видимых одновременно
        1, // номер текущей страницы
        "search.php?&req=test&phrase=1&view=simple&column=def&sort=def&sortmode=ASC&page=" // url страниц
    );
</script>
<table width=100%><tr><td align='left' width=45%></td><td align=center width=10%><font size="3" color="gray"><a href="search.php?&req=test&phrase=1&view=simple&column=def&sort=def&sortmode=ASC&page=2">&nbsp;&nbsp;&#9658;</a></font></td><td align='right' width=45%></td></tr></table></body></html>`
	results := 5
	hashes := parseHashes(response, results)

	if hashes[0] != "2F2DBA2A621B693BB95601C16ED680F8" {
		t.Errorf("got: %s, expected: 2F2DBA2A621B693BB95601C16ED680F8", hashes[0])
	}
	if hashes[1] != "06E6135019C8F2F43158ABA9ABDC610E" {
		t.Errorf("got: %s, expected: 06E6135019C8F2F43158ABA9ABDC610E", hashes[1])
	}
	if hashes[2] != "4363AD191DB6B625BC6200326A51E5DF" {
		t.Errorf("got: %s, expected: 4363AD191DB6B625BC6200326A51E5DF", hashes[2])
	}
	if hashes[3] != "255C181F183AF07A7B8EA95433927DDE" {
		t.Errorf("got: %s, expected: 255C181F183AF07A7B8EA95433927DDE", hashes[3])
	}
	if hashes[4] != "E185937818FE3285AEDDF19B9A85EF9B" {
		t.Errorf("got: %s, expected: E185937818FE3285AEDDF19B9A85EF9B", hashes[4])
	}
}

func TestParseResponse(t *testing.T) {
	// Test on 2F2DBA2A621B693BB95601C16ED680F8
	searchMirror := GetWorkingMirror(SearchMirrors)

	searchMirror.Path = "json.php"
	q := searchMirror.Query()
	q.Set("ids", "2F2DBA2A621B693BB95601C16ED680F8")
	q.Set("fields", JSONQuery)
	searchMirror.RawQuery = q.Encode()

	r, _ := http.Get(searchMirror.String())
	b, _ := ioutil.ReadAll(r.Body)

	_, err := parseResponse(b)
	if err != nil {
		t.Error(err)
	}
}
