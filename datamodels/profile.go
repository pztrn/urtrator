// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package datamodels

type Profile struct {
    // Profile name.
    Name                string          `db:"name"`
    // Game version.
    Version             string          `db:"version"`
    // Binary path.
    Binary              string          `db:"binary"`
    // Will we use second X session?
    Second_x_session    string          `db:"second_x_session"`
    // Additional game parameters we will pass.
    Additional_params   string          `db:"additional_parameters"`
    // Profile path.
    Profile_path        string          `db:"profile_path"`
}
