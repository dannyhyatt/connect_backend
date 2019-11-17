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


UPDATE charity_posts SET thumbnail='https://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg' WHERE thumbnail='http://cdn.sheknows.com/articles/2012/02/Sarah_Parenting/volunteer.jpg';

select title, content, author_id, charity_id, thumbnail, last_edit from (
         select distinct on (post_time) *
         from charity_posts
         order by post_time
     ) t WHERE charity_id in (SELECT charity_id FROM  followers WHERE user_id=1) -- AND post_id NOT IN (SELECT post_id FROM viewed_post WHERE user_id=x) <--- that's untested
     order by post_time limit 10;

SELECT * FROM users;

SELECT id, short_name, long_name, description, profile_url FROM charities WHERE LOWER(short_name) LIKE '%' || LOWER('f') || '%' OR LOWER(long_name) LIKE '%' || LOWER('wf') || '%' LIMIT 5;

SELECT id FROM users WHERE email='daf281@aol.com';

SELECT * FROM followers;
SELECT * FROM charities;