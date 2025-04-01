/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xandelt1@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance_wl

import (
	"github.com/gin-gonic/gin"
)

type AmbulancesAPI interface {


    // CreateAmbulance Post /api/ambulance
    // Saves new ambulance definition 
     CreateAmbulance(c *gin.Context)

    // DeleteAmbulance Delete /api/ambulance/:ambulanceId
    // Deletes specific ambulance 
     DeleteAmbulance(c *gin.Context)

}