// Code generated by go generate; DO NOT EDIT.

package template // import "miniflux.app/template"

var templateCommonMap = map[string]string{
	"entry_pagination": `{{ define "entry_pagination" }}
<div class="pagination">
    <div class="pagination-prev">
        {{ if .prevEntry }}
            <a href="{{ .prevEntryRoute }}{{ if .searchQuery }}?q={{ .searchQuery }}{{ end }}" title="{{ .prevEntry.Title }}" data-page="previous">{{ t "Previous" }}</a>
        {{ else }}
            {{ t "Previous" }}
        {{ end }}
    </div>

    <div class="pagination-next">
        {{ if .nextEntry }}
            <a href="{{ .nextEntryRoute }}{{ if .searchQuery }}?q={{ .searchQuery }}{{ end }}" title="{{ .nextEntry.Title }}" data-page="next">{{ t "Next" }}</a>
        {{ else }}
            {{ t "Next" }}
        {{ end }}
    </div>
</div>
{{ end }}`,
	"item_meta": `{{ define "item_meta" }}
<div class="item-meta">
    <ul>
        <li>
            <a href="{{ route "feedEntries" "feedID" .entry.Feed.ID }}" title="{{ .entry.Feed.SiteURL }}">{{ truncate .entry.Feed.Title 35 }}</a>
        </li>
        <li>
            <time datetime="{{ isodate .entry.Date }}" title="{{ isodate .entry.Date }}">{{ elapsed .user.Timezone .entry.Date }}</time>
        </li>
        {{ if .hasSaveEntry }}
            <li>
                <a href="#"
                    title="{{ t "Save this article" }}"
                    data-save-entry="true"
                    data-save-url="{{ route "saveEntry" "entryID" .entry.ID }}"
                    data-label-loading="{{ t "Saving..." }}"
                    data-label-done="{{ t "Done!" }}"
                    >{{ t "Save" }}</a>
            </li>
        {{ end }}
        <li>
            <a href="{{ .entry.URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ t "Original" }}</a>
        </li>
        {{ if .entry.CommentsURL }}
            <li>
                <a href="{{ .entry.CommentsURL }}" title="{{ t "View Comments" }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer">{{ t "Comments" }}</a>
            </li>
        {{ end }}
        <li>
            <a href="#"
                data-toggle-bookmark="true"
                data-bookmark-url="{{ route "toggleBookmark" "entryID" .entry.ID }}"
                data-label-loading="{{ t "Saving..." }}"
                data-label-star="☆ {{ t "Star" }}"
                data-label-unstar="★ {{ t "Unstar" }}"
                data-value="{{ if .entry.Starred }}star{{ else }}unstar{{ end }}"
                >{{ if .entry.Starred }}★ {{ t "Unstar" }}{{ else }}☆ {{ t "Star" }}{{ end }}</a>
        </li>
        <li>
            <a href="#"
                title="{{ t "Change entry status" }}"
                data-toggle-status="true"
                data-label-read="✔&#xfe0e; {{ t "Read" }}"
                data-label-unread="✘ {{ t "Unread" }}"
                data-value="{{ if eq .entry.Status "read" }}read{{ else }}unread{{ end }}"
                >{{ if eq .entry.Status "read" }}✘ {{ t "Unread" }}{{ else }}✔&#xfe0e; {{ t "Read" }}{{ end }}</a>
        </li>
    </ul>
</div>
{{ end }}`,
	"layout": `{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{template "title" .}} - Miniflux</title>

    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
    <meta name="mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-title" content="Miniflux">
    <link rel="manifest" href="{{ route "webManifest" }}">

    <meta name="robots" content="noindex,nofollow">
    <meta name="referrer" content="no-referrer">

    <!-- Favicons -->
    <link rel="icon" type="image/png" sizes="16x16" href="{{ route "appIcon" "filename" "favicon-16.png" }}">
    <link rel="icon" type="image/png" sizes="32x32" href="{{ route "appIcon" "filename" "favicon-32.png" }}">

    <!-- Android icons -->
    <link rel="icon" type="image/png" sizes="128x128" href="{{ route "appIcon" "filename" "icon-128.png" }}">
    <link rel="icon" type="image/png" sizes="192x192" href="{{ route "appIcon" "filename" "icon-192.png" }}">

    <!-- iOS icons -->
    <link rel="apple-touch-icon" sizes="120x120" href="{{ route "appIcon" "filename" "icon-120.png" }}">
    <link rel="apple-touch-icon" sizes="152x152" href="{{ route "appIcon" "filename" "icon-152.png" }}">
    <link rel="apple-touch-icon" sizes="167x167" href="{{ route "appIcon" "filename" "icon-167.png" }}">
    <link rel="apple-touch-icon" sizes="180x180" href="{{ route "appIcon" "filename" "icon-180.png" }}">

    {{ if .csrf }}
        <meta name="X-CSRF-Token" value="{{ .csrf }}">
    {{ end }}

    <meta name="theme-color" content="{{ theme_color .theme }}">
    <link rel="stylesheet" type="text/css" href="{{ route "stylesheet" "name" .theme }}?{{ .theme_checksum }}">

    <script type="text/javascript" src="{{ route "javascript" "name" "app" }}?{{ .app_js_checksum }}" defer></script>
    <script type="text/javascript" src="{{ route "javascript" "name" "sw" }}?{{ .sw_js_checksum }}" defer id="service-worker-script"></script>
</head>
<body data-entries-status-url="{{ route "updateEntriesStatus" }}">
    {{ if .user }}
    <header class="header">
        <nav>
            <div class="logo">
                <a href="{{ route "unread" }}">Mini<span>flux</span></a>
            </div>
            <ul>
                <li {{ if eq .menu "unread" }}class="active"{{ end }} title="{{ t "Keyboard Shortcut: %s" "g u" }}">
                    <a href="{{ route "unread" }}" data-page="unread">{{ t "Unread" }}
                      {{ if gt .countUnread 0 }}
                          <span class="unread-counter-wrapper">(<span class="unread-counter">{{ .countUnread }}</span>)</span>
                      {{ end }}
                    </a>
                </li>
                <li {{ if eq .menu "starred" }}class="active"{{ end }} title="{{ t "Keyboard Shortcut: %s" "g b" }}">
                    <a href="{{ route "starred" }}" data-page="starred">{{ t "Starred" }}</a>
                </li>
                <li {{ if eq .menu "history" }}class="active"{{ end }} title="{{ t "Keyboard Shortcut: %s" "g h" }}">
                    <a href="{{ route "history" }}" data-page="history">{{ t "History" }}</a>
                </li>
                <li {{ if eq .menu "feeds" }}class="active"{{ end }} title="{{ t "Keyboard Shortcut: %s" "g f" }}">
                    <a href="{{ route "feeds" }}" data-page="feeds">{{ t "Feeds" }}
                      {{ if gt .countErrorFeeds 0 }}
                          <span class="error-feeds-counter-wrapper">(<span class="error-feeds-counter">{{ .countErrorFeeds }}</span>)</span>
                      {{ end }}
                    </a>
                </li>
                <li {{ if eq .menu "categories" }}class="active"{{ end }} title="{{ t "Keyboard Shortcut: %s" "g c" }}">
                    <a href="{{ route "categories" }}" data-page="categories">{{ t "Categories" }}</a>
                </li>
                <li {{ if eq .menu "settings" }}class="active"{{ end }} title="{{ t "Keyboard Shortcut: %s" "g s" }}">
                    <a href="{{ route "settings" }}" data-page="settings">{{ t "Settings" }}</a>
                </li>
                <li>
                    <a href="{{ route "logout" }}" title="{{ t "Logged as %s" .user.Username }}">{{ t "Logout" }}</a>
                </li>
            </ul>
            <div class="search">
                <div class="search-toggle-switch {{ if $.searchQuery }}has-search-query{{ end }}">
                    <a href="#" data-action="search">&laquo;&nbsp;{{ t "Search" }}</a>
                </div>
                <form action="{{ route "searchEntries" }}" class="search-form {{ if $.searchQuery }}has-search-query{{ end }}">
                    <input type="search" name="q" id="search-input" placeholder="{{ t "Search..." }}" {{ if $.searchQuery }}value="{{ .searchQuery }}"{{ end }} required>
                </form>
            </div>
        </nav>
    </header>
    {{ end }}
    {{ if .flashMessage }}
        <div class="flash-message alert alert-success">{{ .flashMessage }}</div>
    {{ end }}
    {{ if .flashErrorMessage }}
        <div class="flash-error-message alert alert-error">{{ .flashErrorMessage }}</div>
    {{ end }}
    <main>
        {{template "content" .}}
    </main>
    <template id="keyboard-shortcuts">
        <div id="modal-left">
            <a href="#" class="btn-close-modal">x</a>
            <h3>{{ t "Keyboard Shortcuts" }}</h3>

            <div class="keyboard-shortcuts">
                <p>{{ t "Sections Navigation" }}</p>
                <ul>
                    <li>{{ t "Go to unread" }} = <strong>g + u</strong></li>
                    <li>{{ t "Go to bookmarks" }} = <strong>g + b</strong></li>
                    <li>{{ t "Go to history" }} = <strong>g + h</strong></li>
                    <li>{{ t "Go to feeds" }} = <strong>g + f</strong></li>
                    <li>{{ t "Go to categories" }} = <strong>g + c</strong></li>
                    <li>{{ t "Go to settings" }} = <strong>g + s</strong></li>
                    <li>{{ t "Show keyboard shortcuts" }} = <strong>?</strong></li>
                </ul>

                <p>{{ t "Items Navigation" }}</p>
                <ul>
                    <li>{{ t "Go to previous item" }} = <strong>p {{ t "or" }} j {{ t "or" }} ◄</strong></li>
                    <li>{{ t "Go to next item" }} = <strong>n {{ t "or" }} k {{ t "or" }} ►</strong></li>
                </ul>

                <p>{{ t "Pages Navigation" }}</p>
                <ul>
                    <li>{{ t "Go to previous page" }} = <strong>h</strong></li>
                    <li>{{ t "Go to next page" }} = <strong>l</strong></li>
                </ul>

                <p>{{ t "Actions" }}</p>
                <ul>
                    <li>{{ t "Open selected item" }} = <strong>o</strong></li>
                    <li>{{ t "Open original link" }} = <strong>v</strong></li>
                    <li>{{ t "Toggle read/unread" }} = <strong>m</strong></li>
                    <li>{{ t "Mark current page as read" }} = <strong>A</strong></li>
                    <li>{{ t "Download original content" }} = <strong>d</strong></li>
                    <li>{{ t "Toggle bookmark" }} = <strong>f</strong></li>
                    <li>{{ t "Save article" }} = <strong>s</strong></li>
                    <li>{{ t "Set focus on search form" }} = <strong>/</strong></li>
                    <li>{{ t "Close modal dialog" }} = <strong>Esc</strong></li>
                </ul>
            </div>
        </div>
    </template>
</body>
</html>
{{ end }}
`,
	"pagination": `{{ define "pagination" }}
<div class="pagination">
    <div class="pagination-prev">
        {{ if .ShowPrev }}
            <a href="{{ .Route }}{{ if gt .PrevOffset 0 }}?offset={{ .PrevOffset }}{{ if .SearchQuery }}&amp;q={{ .SearchQuery }}{{ end }}{{ else }}{{ if .SearchQuery }}?q={{ .SearchQuery }}{{ end }}{{ end }}" data-page="previous">{{ t "Previous" }}</a>
        {{ else }}
            {{ t "Previous" }}
        {{ end }}
    </div>

    <div class="pagination-next">
        {{ if .ShowNext }}
            <a href="{{ .Route }}?offset={{ .NextOffset }}{{ if .SearchQuery }}&amp;q={{ .SearchQuery }}{{ end }}" data-page="next">{{ t "Next" }}</a>
        {{ else }}
            {{ t "Next" }}
        {{ end }}
    </div>
</div>
{{ end }}
`,
}

var templateCommonMapChecksums = map[string]string{
	"entry_pagination": "756ef122f3ebc73754b5fc4304bf05e59da0ab4af030b2509ff4c9b4a74096ce",
	"item_meta":        "d85d1ae181120b7a3d38173b3990a0c2a5cf706cabdfb00c4d002a09271de51e",
	"layout":           "2491695e33a496c9bd902a2cb5bc3a6a540f98ac7c24591d503a77ba0f5f0ebe",
	"pagination":       "b592d58ea9d6abf2dc0b158621404cbfaeea5413b1c8b8b9818725963096b196",
}
