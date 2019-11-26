insert into charities(short_name, long_name, description, total_donated, queued_donations, ceo, profile_url, password_hash)
 values (
         'PETA',
         'Pets Eels Tails Association',
         'We do our best to stop animal cruelty around the world.',
         233343, -- cents
         11293,
         'Brooke Hammond',
         'https://shop.peta.org/media/catalog/product/cache/ecd051e9670bd57df35c8f0b122d8aea/p/t/pta10072.jpg',
         'petapassword'
        );

insert into charities(short_name, long_name, description, total_donated, queued_donations, ceo, profile_url, password_hash)
 values (
         'WWF',
         'World Wildlife Fund',
         'We help prevent loss of habitat and help animals all around the globe.',
         24233343, -- cents
         1124293,
         'Suprete Bjord',
         'https://media.kidozi.com/unsafe/150x150/img.kidozi.com/design/150/150/ffffff/47916/a3202e30dfcd7bf8249171024f34d98f.png.jpg',
         'wwfpassword'
        );

insert into charities(short_name, long_name, description, total_donated, queued_donations, ceo, profile_url, password_hash)
 values (
         'TBF',
         'That Bird Foundation',
         'We go to random houses and save birds.',
         23434, -- cents
         123244,
         'Neighborhood Guy',
         'https://media.kidozi.com/unsafe/150x150/img.kidozi.com/design/150/150/ffffff/47916/a3202e30dfcd7bf8249171024f34d98f.png.jpg',
         'tbfpassword'
        );

select * from charities;


insert into charity_users(charity_id, display_name, bio, password_hash)
 VALUES (
         1,
         'Johansen Alberta',
         'I work for PeTA!',
         'petaposter'
        );

insert into charity_users(charity_id, display_name, bio, password_hash)
 VALUES (
         2,
         'Sanjay McCunnon',
         'I work for the WWF!',
         'wwfposter'
        );

insert into charity_users(charity_id, display_name, bio, password_hash)
 VALUES (
         3,
         'Door ToDoor Guy',
         'I work for Thaat Bird Foundation',
         'tbfposter'
        );


select * from charity_users;

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         1,
         1,
         'PeTa is stopping animal cruelty!',
         '# How peta is stoping animal cruelty \nayayyaya\nfrick the farmers',
         'https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         1,
         1,
         'PeTa post 2!',
         '# How peta is stoping animal cruelty \nayayyaya\nfrick the farmers',
         'https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg',
         now(),
         now()
        );
insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         1,
         1,
         'PeTa post 3',
         '# How peta is stoping animal cruelty \nayayyaya\nfrick the farmers',
         'https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         1,
         1,
         'PeTa post 4',
         '# How peta is stoping animal cruelty \nayayyaya\nfrick the farmers',
         'https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         1,
         1,
         'PeTa post 5!',
         '# How peta is stoping animal cruelty \nayayyaya\nfrick the farmers',
         'https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         2,
         2,
         'WWF is saving the turtles',
         '# dont use straws \neeek\nvsco the the the farmers',
         'https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         2,
         2,
         'WWF is saving the toes',
         '# dont use straws fgterfwdwrgtgfwg dgssd h htddddh d hd fd  hrd ht  hhdd h d hnvsco the the the farmers',
         'https://s3.amazonaws.com/cn-web-site/landingpages/love_a_charity_lp_img.jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         3,
         3,
         'TBF is saving the Birds!',
         '# dont kill the birds\n we go around to peoples doors to see if they need help',
         'https://www.want2donate.org/sites/default/files/imag960x560rspb9-0..jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         3,
         3,
         'TBF post 2!',
         '# dont kill the birds\n we go around to peoples doors to see if they need help',
         'https://www.want2donate.org/sites/default/files/imag960x560rspb9-0..jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         3,
         3,
         'TBF post 3!',
         '# dont kill the birds\n we go around to peoples doors to see if they need help',
         'https://www.want2donate.org/sites/default/files/imag960x560rspb9-0..jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         3,
         3,
         'TBF post 4!',
         '# dont kill the birds\n we go around to peoples doors to see if they need help',
         'https://www.want2donate.org/sites/default/files/imag960x560rspb9-0..jpg',
         now(),
         now()
        );

insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         3,
         3,
         'TBF post 5!',
         '# dont kill the birds\n we go around to peoples doors to see if they need help',
         'https://www.want2donate.org/sites/default/files/imag960x560rspb9-0..jpg',
         now(),
         now()
        );
insert into charity_posts(charity_id, author_id, title, content, thumbnail, post_time, last_edit)
 VALUES (
         3,
         3,
         'TBF post sIX',
         '# Markdown: Syntax

*   [Overview](#overview)
    *   [Philosophy](#philosophy)
    *   [Inline HTML](#html)
    *   [Automatic Escaping for Special Characters](#autoescape)
*   [Block Elements](#block)
    *   [Paragraphs and Line Breaks](#p)
    *   [Headers](#header)
    *   [Blockquotes](#blockquote)
    *   [Lists](#list)
    *   [Code Blocks](#precode)
    *   [Horizontal Rules](#hr)
*   [Span Elements](#span)
    *   [Links](#link)
    *   [Emphasis](#em)
    *   [Code](#code)
    *   [Images](#img)
*   [Miscellaneous](#misc)
    *   [Backslash Escapes](#backslash)
    *   [Automatic Links](#autolink)


**Note:** This document is itself written using Markdown; you
can [see the source for it by adding ''.text'' to the URL](/projects/markdown/syntax.text).

----

## Overview

### Philosophy

Markdown is intended to be as easy-to-read and easy-to-write as is feasible.

Readability, however, is emphasized above all else. A Markdown-formatted
document should be publishable as-is, as plain text, without looking
like it''s been marked up with tags or formatting instructions. While
Markdown''s syntax has been influenced by several existing text-to-HTML
filters -- including [Setext](http://docutils.sourceforge.net/mirror/setext.html), [atx](http://www.aaronsw.com/2002/atx/), [Textile](http://textism.com/tools/textile/), [reStructuredText](http://docutils.sourceforge.net/rst.html),
[Grutatext](http://www.triptico.com/software/grutatxt.html), and [EtText](http://ettext.taint.org/doc/) -- the single biggest source of
inspiration for Markdown''s syntax is the format of plain text email.

## Block Elements

### Paragraphs and Line Breaks

A paragraph is simply one or more consecutive lines of text, separated
by one or more blank lines. (A blank line is any line that looks like a
blank line -- a line containing nothing but spaces or tabs is considered
blank.) Normal paragraphs should not be indented with spaces or tabs.

The implication of the "one or more consecutive lines of text" rule is
that Markdown supports "hard-wrapped" text paragraphs. This differs
significantly from most other text-to-HTML formatters (including Movable
Type''s "Convert Line Breaks" option) which translate every line break
character in a paragraph into a `<br />` tag.

When you *do* want to insert a `<br />` break tag using Markdown, you
end a line with two or more spaces, then type return.

### Headers

Markdown supports two styles of headers, [Setext] [1] and [atx] [2].

Optionally, you may "close" atx-style headers. This is purely
cosmetic -- you can use this if you think it looks better. The
closing hashes don''t even need to match the number of hashes
used to open the header. (The number of opening hashes
determines the header level.)


### Blockquotes

Markdown uses email-style `>` characters for blockquoting. If you''re
familiar with quoting passages of text in an email message, then you
know how to create a blockquote in Markdown. It looks best if you hard
wrap the text and put a `>` before every line:

> This is a blockquote with two paragraphs. Lorem ipsum dolor sit amet,
> consectetuer adipiscing elit. Aliquam hendrerit mi posuere lectus.
> Vestibulum enim wisi, viverra nec, fringilla in, laoreet vitae, risus.
>
> Donec sit amet nisl. Aliquam semper ipsum sit amet velit. Suspendisse
> id sem consectetuer libero luctus adipiscing.

Markdown allows you to be lazy and only put the `>` before the first
line of a hard-wrapped paragraph:

> This is a blockquote with two paragraphs. Lorem ipsum dolor sit amet,
consectetuer adipiscing elit. Aliquam hendrerit mi posuere lectus.
Vestibulum enim wisi, viverra nec, fringilla in, laoreet vitae, risus.

> Donec sit amet nisl. Aliquam semper ipsum sit amet velit. Suspendisse
id sem consectetuer libero luctus adipiscing.

Blockquotes can be nested (i.e. a blockquote-in-a-blockquote) by
adding additional levels of `>`:

> This is the first level of quoting.
>
> > This is nested blockquote.
>
> Back to the first level.

Blockquotes can contain other Markdown elements, including headers, lists,
and code blocks:

> ## This is a header.
>
> 1.   This is the first list item.
> 2.   This is the second list item.
>
> Here''s some example code:
>
>     return shell_exec("echo $input | $markdown_script");

Any decent text editor should make email-style quoting easy. For
example, with BBEdit, you can make a selection and choose Increase
Quote Level from the Text menu.


### Lists

Markdown supports ordered (numbered) and unordered (bulleted) lists.

Unordered lists use asterisks, pluses, and hyphens -- interchangably
-- as list markers:

*   Red
*   Green
*   Blue

is equivalent to:

+   Red
+   Green
+   Blue

and:

-   Red
-   Green
-   Blue

Ordered lists use numbers followed by periods:

1.  Bird
2.  McHale
3.  Parish

It''s important to note that the actual numbers you use to mark the
list have no effect on the HTML output Markdown produces. The HTML
Markdown produces from the above list is:

If you instead wrote the list in Markdown like this:

1.  Bird
1.  McHale
1.  Parish

or even:

3. Bird
1. McHale
8. Parish

you''d get the exact same HTML output. The point is, if you want to,
you can use ordinal numbers in your ordered Markdown lists, so that
the numbers in your source match the numbers in your published HTML.
But if you want to be lazy, you don''t have to.

To make lists look nice, you can wrap items with hanging indents:

*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
    Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
    viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
    Suspendisse id sem consectetuer libero luctus adipiscing.

But if you want to be lazy, you don''t have to:

*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
Suspendisse id sem consectetuer libero luctus adipiscing.

List items may consist of multiple paragraphs. Each subsequent
paragraph in a list item must be indented by either 4 spaces
or one tab:

1.  This is a list item with two paragraphs. Lorem ipsum dolor
    sit amet, consectetuer adipiscing elit. Aliquam hendrerit
    mi posuere lectus.

    Vestibulum enim wisi, viverra nec, fringilla in, laoreet
    vitae, risus. Donec sit amet nisl. Aliquam semper ipsum
    sit amet velit.

2.  Suspendisse id sem consectetuer libero luctus adipiscing.

It looks nice if you indent every line of the subsequent
paragraphs, but here again, Markdown will allow you to be
lazy:

*   This is a list item with two paragraphs.

    This is the second paragraph in the list item. You''re
only required to indent the first line. Lorem ipsum dolor
sit amet, consectetuer adipiscing elit.

*   Another item in the same list.

To put a blockquote within a list item, the blockquote''s `>`
delimiters need to be indented:

*   A list item with a blockquote:

    > This is a blockquote
    > inside a list item.

To put a code block within a list item, the code block needs
to be indented *twice* -- 8 spaces or two tabs:

*   A list item with a code block:

        <code goes here>

### Code Blocks

Pre-formatted code blocks are used for writing about programming or
markup source code. Rather than forming normal paragraphs, the lines
of a code block are interpreted literally. Markdown wraps a code block
in both `<pre>` and `<code>` tags.

To produce a code block in Markdown, simply indent every line of the
block by at least 4 spaces or 1 tab.

This is a normal paragraph:

    This is a code block.

Here is an example of AppleScript:

    tell application "Foo"
        beep
    end tell

A code block continues until it reaches a line that is not indented
(or the end of the article).

Within a code block, ampersands (`&`) and angle brackets (`<` and `>`)
are automatically converted into HTML entities. This makes it very
easy to include example HTML source code using Markdown -- just paste
it and indent it, and Markdown will handle the hassle of encoding the
ampersands and angle brackets. For example, this:

    <div class="footer">
        &copy; 2004 Foo Corporation
    </div>

Regular Markdown syntax is not processed within code blocks. E.g.,
asterisks are just literal asterisks within a code block. This means
it''s also easy to use Markdown to write about Markdown''s own syntax.

```
tell application "Foo"
    beep
end tell
```

## Span Elements

### Links

Markdown supports two style of links: *inline* and *reference*.

In both styles, the link text is delimited by [square brackets].

To create an inline link, use a set of regular parentheses immediately
after the link text''s closing square bracket. Inside the parentheses,
put the URL where you want the link to point, along with an *optional*
title for the link, surrounded in quotes. For example:

This is [an example](http://example.com/) inline link.

[This link](http://example.net/) has no title attribute.

### Emphasis

Markdown treats asterisks (`*`) and underscores (`_`) as indicators of
emphasis. Text wrapped with one `*` or `_` will be wrapped with an
HTML `<em>` tag; double `*`''s or `_`''s will be wrapped with an HTML
`<strong>` tag. E.g., this input:

*single asterisks*

_single underscores_

**double asterisks**

__double underscores__

### Code

To indicate a span of code, wrap it with backtick quotes (`` ` ``).
Unlike a pre-formatted code block, a code span indicates code within a
normal paragraph. For example:

Use the `printf()` function.
',
         'https://www.want2donate.org/sites/default/files/imag960x560rspb9-0..jpg',
         now(),
         now()
        );
UPDATE charity_posts SET thumbnail='https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg' WHERE thumbnail='http://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg';



select * from charity_posts;


insert into followers(user_id, charity_id)
 VALUES (
         1, 1
        );

insert into followers(user_id, charity_id)
 VALUES (
         1, 2
        );

select * from followers where user_id=1 and charity_id=2;

SELECT * FROM charity_posts WHERE id=1;

INSERT INTO views (user_id, charity_post_id, viewed_at) VALUES (1, unnest(array(
select id from (
         select distinct on (last_edit) *
         from charity_posts
         order by last_edit desc
     ) t WHERE charity_id in (SELECT charity_id FROM  followers WHERE user_id=1) AND id NOT IN (SELECT charity_post_id FROM views WHERE user_id=1) -- that's untested
     order by last_edit limit 5)
), now());

SELECT * FROM charity_posts REVERSE ORDER BY last_edit DESC;

SELECT * FROM views WHERE user_id=1 ORDER BY viewed_at DESC LIMIT 5;
DELETE FROM views WHERE user_id=1;

SELECT * FROM views;

SELECT email, name, phone_number FROM users WHERE id=1;
--SELECT SUM(SELECT amount FROM donations WHERE id=1 AND flushed=true);

SELECT id, short_name, long_name, description, profile_url FROM charities WHERE LOWER(short_name) LIKE '%' || LOWER('f') || '%' OR LOWER(long_name) LIKE '%' || LOWER('wf') || '%' LIMIT 5;

SELECT id FROM users WHERE email='daf281@aol.com';

SELECT * FROM followers;
SELECT * FROM charities;
SELECT  * FROM charity_posts;