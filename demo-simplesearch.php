<?php
/**
 * demo.php - read the files.json, inverted-wordlist.json file and
 * offer a naive search implementation by find the interection of
 * of words in the request.
 */
if (php_sapi_name() !== 'cli') {
    die('demo.php is a command line tool.');
}
$files = json_decode(file_get_contents('files.json'), true);
$words = json_decode(file_get_contents('inverted-wordlist.json'), true);

if (count($argv) < 2) {
    die('USAGE php demo.php SEARCH_WORDS');
}

$pages = [];
$found_words = [];
for ($i = 1; $i < count($argv); $i += 1) {
    $term = strtolower($argv[$i]);
    if (isset($words[$term]))  {
        $pages = array_merge($pages, $words[$term]);
        $found_words[] = $term;
    }
}
echo "These are the pages that have some of the words (i.e. " . implode($found_words, ", ") . "):\n";
foreach ($pages as $entry) {
    echo "Page: " . $files[$entry] . PHP_EOL;
}
?>
