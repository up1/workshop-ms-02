package service.profile.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.sun.org.apache.xerces.internal.xs.StringList;
import org.springframework.stereotype.Service;
import service.profile.entity.Feed;
import service.profile.entity.Like;
import service.profile.entity.Profile;

import javax.annotation.PostConstruct;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@Service
public class ProfileService {


    public Profile getProfile(int id) {
        Profile profile = new Profile();
        profile.setId(id);
        profile.setName("Demo user" + (id + 1));
        profile.setFeeds(getFeelds(id));

        return profile;
    }

    public List<Feed> getFeelds(int profileId) {
        List<Feed> feeds = new ArrayList<>();
        for (int i = 0; i < 10; i++) {
            Like like = new Like();
            like.setCount((i+1) * 10);
            like.setName(new String[] {
                    "name1",
                    "name2",
                    "name3"
            });

            Feed feed = new Feed();
            feed.setId(profileId + 1);
            feed.setDescription("feed item " + (i + 1));
            feed.setLike(like);
            feeds.add(feed);
        }

        return feeds;
    }
}
