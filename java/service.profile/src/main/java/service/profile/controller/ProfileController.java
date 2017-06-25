package service.profile.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestParam;
import service.profile.entity.Profile;
import service.profile.service.ProfileService;
import org.springframework.beans.factory.annotation.Autowired;

@RestController
public class ProfileController{

    @Autowired
    ProfileService profileService;

    @GetMapping("/profile/{id}")
    public Profile getProfile(@PathVariable(value="id") int id){

        Profile profile = profileService.getProfile(1);
        return profile;
    }
}
