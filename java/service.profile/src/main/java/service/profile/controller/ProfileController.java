package service.profile;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestParam;

@RestController
public class ProfileController{

    @GetMapping("/profile")
    public String getProfile(@RequestParam(value="ID", defaultValue="") String ID){

        if (ID.equals("")) return "service require ID";

        return ID;
    }

}
