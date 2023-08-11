use test;

# 目前的数据可以支持users和videos的映射
INSERT INTO users (name, follow_count, follower_count, is_follow, avatar, background_image, signature, total_favorited, work_count, favorite_count)
VALUES
    ('John Doe', 150, 200, true, 'avatar1.jpg', 'bg1.jpg', 'Hello, I am John Doe', 500, 50, 100),
    ('Jane Smith', 200, 150, false, 'avatar2.jpg', 'bg2.jpg', 'Greetings from Jane Smith', 800, 70, 120),
    ('Michael Johnson', 100, 120, true, 'avatar3.jpg', 'bg3.jpg', 'Hi, I am Michael Johnson', 300, 40, 80);

INSERT INTO videos (user_id, play_url, cover_url, favorite_count, comment_count, is_favorite, title)
VALUES
    (1, 'https://example.com/video1.mp4', 'https://example.com/cover1.jpg', 50, 30, true, 'Video 1'),
    (2, 'https://example.com/video2.mp4', 'https://example.com/cover2.jpg', 80, 40, false, 'Video 2'),
    (3, 'https://example.com/video3.mp4', 'https://example.com/cover3.jpg', 120, 20, true, 'Video 3');

-- John Doe 喜欢 Video 2
INSERT INTO user_like_videos (user_id, video_id) VALUES (1, 2);
-- Jane Smith 喜欢 Video 1 和 Video 3
INSERT INTO user_like_videos (user_id, video_id) VALUES (2, 1), (2, 3);
-- Michael Johnson 喜欢 Video 1
INSERT INTO user_like_videos (user_id, video_id) VALUES (3, 1);
