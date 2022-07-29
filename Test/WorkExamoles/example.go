package main

// модифікація запису на LDAP

/*dev.GET("/modify-branch-net", func(c echo.Context) error {
	serial := "4AC7026F71B4"

	//model, err := devicemodel.Instance()
	//if err != nil {
	//	return err
	//}
	//device, err := model.GetBySerial(serial)
	//if err != nil {
	//	return err
	//}
	//
	//logger.Debug("ban %+v", device)

	ldapClient, err := ldapprovider.New()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	data, err := ldapClient.Find(serial, ldapprovider.DefaultAttrs)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	data["radiusFramedRoute"] = []string{"10.31.71.128/26", "10.31.103.128/26"}

	err = ldapClient.Modify(serial, data)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, data)
})*/

//OldNetwork  ProfileBranchNetwork `json:"old_network" bson:"old_network"`
