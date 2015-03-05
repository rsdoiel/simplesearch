/**
 * words_test.go - tests for words.go
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package words

import (
	"bytes"
	"strings"
	"testing"
)

func TestWords(t *testing.T) {
	w := new(Words)
	if w == nil {
		t.Error("Could not create Words struct.")
	}
	w.Words = make(map[string][]int, 100)
	w.Words["one"] = append(w.Words["one"], 1)
	if w.Words["one"][0] != 1 {
		t.Error("Could not append 1 to Words.")
	}

	w2 := New()
	if w2 == nil {
		t.Error(".New() failed to create a new Words structure")
	}
	w2.Files = append(w2.Files, "test-01.txt")
	w2.Words["Hello"] = append(w2.Words["Hello"], 0)
	w2.Words["World"] = append(w2.Words["World"], 0)

	if len(w2.Files) != 1 {
		t.Error(".Files should be length 1")
	}
	if len(w2.Words["Hello"]) != 1 {
		t.Error(".Words[\"Hello\"] should be length 1")
	}
	if len(w2.Words["World"]) != 1 {
		t.Error(".Words[\"World\"] should be length 1")
	}
	if w2.Words["Fred"] != nil {
		t.Error(".Words[\"Fred\"] should be nil")
	}
}

func TestFlatten(t *testing.T) {
	expectedText := []byte("This is a paragraph")
	plainText := Flatten([]byte("<p>This is a paragraph</p>"))
	if bytes.Equal(expectedText, plainText) != true {
		t.Errorf("[%s] != [%s]\n", expectedText, plainText)
	}
	expectedText = []byte("This is a test but without the script element")
	plainText = Flatten([]byte("<body><p>This is a test<script>console.log('Should not include');</script>, but without the script element.</body>"))
	if len(expectedText) != len(plainText) {
		t.Errorf("different lengths %d <> %d\n", len(expectedText), len(plainText))
	}
	if bytes.Equal(expectedText, plainText) == false {
		t.Errorf("[%s] != [%s]\n", expectedText, plainText)
	}
}

func TestsWordList(t *testing.T) {
	src := []byte("<body><header>This si a test</header><h1>A title</h1><p>and a paragraph</p></body>")
	expectedWords := [][]byte{[]byte("This"), []byte("si"), []byte("a"), []byte("test"), []byte("A"), []byte("title"), []byte("and"), []byte("paragraph")}
	s := Flatten(src)
	w := WordList(s)
	if w == nil {
		t.Errorf("WordList() returned empty list")
	}
	if len(w) != len(expectedWords) {
		t.Errorf("len(w) %d does not equal len(expectedWords) %d\n", len(w), len(expectedWords))
	}
	for _, item := range expectedWords {
		if indexOf(w, string(item[:])) == -1 {
			t.Error("Could not find %v in %v\n", item, w)
		}
	}
}

func TestMergeWords(t *testing.T) {
	w := New()
	if w.MergeWords("test-01.html", [][]byte{[]byte("This"), []byte("iS"), []byte("a"), []byte("TEST")}) == false {
		t.Errorf("MergeWords returned false %v\n", w)
	}
	if len(w.Files) != 1 {
		t.Errorf("Should have a single file in .Files: %v\n", w)
	}
	if containsString(w.Files, "test-01.html") == false {
		t.Errorf("MergeWords() failed did not add words for test.html: %v\n", w)
	}
	if w.Words["this"] == nil && w.Words["this"][0] == 0 {
		t.Errorf("Word 'this' should be in w.Words: %v\n", w.Words)
	}
	if w.Words["is"] == nil && w.Words["is"][0] == 0 {
		t.Errorf("Word 'is' should be in w.Words: %v\n", w.Words)
	}
	if w.Words["a"] == nil && w.Words["a"][0] == 0 {
		t.Errorf("Word 'a' should be in w.Words: %v\n", w.Words)
	}
	if w.Words["test"] == nil && w.Words["test"][0] == 0 {
		t.Errorf("Word 'test' should be in w.Words: %v\n", w.Words)
	}

	if w.MergeWords("test-02.html", [][]byte{[]byte("This"), []byte("is"), []byte("not"), []byte("a"), []byte("unique"), []byte("test")}) == false {
		t.Errorf("MergeWords() returned false %v\n", w)
	}
	if len(w.Files) != 2 {
		t.Errorf("Should have two files in struct: %v\n", w)
	}
	if len(w.Words["this"]) != 2 {
		t.Errorf("Should have two entries for 'this' %v\n", w.Words)
	}
	if len(w.Words["not"]) != 1 {
		t.Errorf("Should have one entry for 'not' %v\n", w.Words)
	}

	if w.MergeWords("test-03.html", [][]byte{[]byte("THIS"), []byte("SPACE"), []byte(""), []byte("?"), []byte("!"), []byte("A"), []byte("\n\r")}) == false {
		t.Errorf("MergeWords() returned false %v\n", w)
	}
	if w.Words[""] != nil {
		t.Errorf("MergeWords() is storing empty string: %v\n", w)
	}
	if w.Words["?"] != nil {
		t.Errorf("MergeWords() is storing question mark: %v\n", w)
	}
	if w.Words["!"] != nil {
		t.Errorf("MergeWords() is storing bang: %v\n", w)
	}
	if w.Words["\n\r"] != nil {
		t.Errorf("MergeWords() is storing line feed/carriage return: %v\n", w)
	}

}

func TestToJSON(t *testing.T) {
	w := New()
	if w.MergeWords("test-01.html", [][]byte{[]byte("This"), []byte("iS"), []byte("a"), []byte("TEST")}) == false {
		t.Errorf("MergeWords returned error %v\n", w)
	}
	if w.MergeWords("test-02.html", [][]byte{[]byte("This"), []byte("is"), []byte("not"), []byte("a"), []byte("unique"), []byte("test")}) == false {
		t.Errorf("MergeWords() returned error %v\n", w)
	}
	fileList, wordList, err := w.ToJSON()
	if err != nil {
		t.Errorf("ToJSON() failed: %v, %v\n", w, err)
	}
	if strings.Contains(fileList, "test-01.html") == false || strings.Contains(fileList, "test-02.html") == false {
		t.Errorf("Missing filenames in list: %v -> %s\n", w.Files, fileList)
	}
	if strings.Contains(wordList, "\"this\":[0,1]") == false {
		t.Errorf("Missing 'This' values in word list: %v -> %s\n", w.Words, wordList)
	}
}

func TestMoreCompletedHTMLProcessing(t *testing.T) {
	page_source := []byte(`
<!DOCTYPE html>
    <html>
    <head>
        <title>Persona demo using the ottoengine</title>
        <meta http-equiv="X-UA-Compatible" content="IE=Edge">
        <script src="https://login.persona.org/include.js"></script>
    </head>
    <body>
    <header><button id="signin">Sign in</button> <button id="signout">Signout</button></header>
        <h1>Persona Demo</h1>
        <caption>Uses <em>ws</em> ottoengine</caption>
        <p>Authenticated with Mozilla's Persona and via ottoengine</p>

<script>
(function (){
    "use strict";
    var signinLink = document.getElementById('signin'),
        signoutLink = document.getElementById('signout');

    if (signinLink) {
        signinLink.onclick = function() { navigator.id.request(); };
    }
    if (signoutLink) {
        signoutLink.onclick = function() { navigator.id.logout(); };
    }

    function simpleXhrSentinel(xhr) {
        return function() {
            if (xhr.readyState == 4) {
                if (xhr.status == 200){
                    // reload page to reflect new login state
                    window.location.reload();
                } else {
                    navigator.id.logout();
                    alert("XMLHttpRequest error: " + xhr.status); 
                } 
            } 
        } 
    }

    function verifyAssertion(assertion) {
        // Your backend must return HTTP status code 200 to indicate successful
        // verification of user's email address and it must arrange for the binding
        // of currentUser to said address when the page is reloaded
        var xhr = new XMLHttpRequest();

        xhr.open("POST", "/persona", true);
        // see http://www.openjs.com/articles/ajax_xmlhttp_using_post.php
        var param = "assertion="+assertion;
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
       xhr.setRequestHeader("Content-length", param.length);
       xhr.setRequestHeader("Connection", "close");
       xhr.send(param); // for verification by your backend
       xhr.onreadystatechange = simpleXhrSentinel(xhr); 
    }

    function signoutUser() {
        // Your backend must return HTTP status code 200 to indicate successful
        // sign out (usually the resetting of one or more session variables) and
        // it must arrange for the binding of currentUser to 'null' when the page
        // is reloaded
        var xhr = new XMLHttpRequest();
        xhr.open("GET", "/xhr/sign-out", true);
        xhr.send(null);
        xhr.onreadystatechange = simpleXhrSentinel(xhr); 
    }
    // Go!
    navigator.id.watch({
        loggedInUser: currentUser,
        onlogin: verifyAssertion,
        onlogout: signoutUser 
    });
}());
</script>
    </body>
</html>`)

	expected_source := []byte("Sign in Signout Persona Demo Uses ws ottoengine Authenticated with Mozilla's Persona and via ottoengine")

	if text := Flatten(page_source); bytes.Equal(expected_source, text) == false {
		t.Errorf("Flatten() failed: %s\n", text)
	}
}

func TestWordListHandling(t *testing.T) {
	flattenedText := []byte("Sign in Signout Persona Demo Uses ws ottoengine Authenticated with Mozilla's Persona and via ottoengine")
	wordlist := WordList(flattenedText)

	if len(wordlist) != 15 {
		t.Errorf("Expected 15 words: %d\n", len(wordlist))
	}
}
