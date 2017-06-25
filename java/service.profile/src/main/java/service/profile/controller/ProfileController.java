package service.profile;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestParam;
import service.profile.entity.Profile;
import service.profile.service.ProfileService;
import org.springframework.beans.factory.annotation.Autowired;

@RestController
public class ProfileController{

    @Autowired
    ProfileService profileService;

    @GetMapping("/profile")
    public Profile getProfile(@RequestParam(value="ID", defaultValue="") String ID){

        if (ID.equals("")) return new Profile();       

        Profile profile = profileService.getProfile(1);
        return profile;
    }
}
