#pragma once

#define ALT_SERVER_API
#ifndef GO_MODULE_VERSION
    #define GO_MODULE_VERSION "DEBUG"
#endif

typedef struct pos {
    float x;
    float y;
    float z;
} Position;

typedef struct rot {
    float roll;
    float pitch;
    float yaw;
} Rotation;

typedef struct rgba {
    unsigned char r;
    unsigned char g;
    unsigned char b;
    unsigned char a;
} RGBA;

typedef struct cloth {
    unsigned int drawableId;
    unsigned int textureId;
    unsigned int paletteId;
} Cloth;

typedef struct dlcCloth {
    unsigned long dlc;
    unsigned int drawableId;
    unsigned int textureId;
    unsigned int paletteId;
} DlcCloth;

typedef struct prop {
    unsigned int drawableId;
    unsigned int textureId;
} Prop;

typedef struct dlcProp {
    unsigned long dlc;
    unsigned int drawableId;
    unsigned int textureId;
} DlcProp;

typedef struct fireInfo {
    Position position;
    unsigned long weaponHash;
} FireInfo;

typedef struct neonState {
    unsigned char left;
    unsigned char right;
    unsigned char front;
    unsigned char back;
} VehicleNeonState;

typedef struct entity {
    unsigned char typ;
    void *ptr;
    unsigned int id;
    unsigned int model;
} Entity;

typedef struct array {
    unsigned long long size;
    void *array;
} Array;

typedef struct connectionInfo {
    unsigned int id;
    const char *name;
    unsigned long long socialID;
    const char *socialName;
    unsigned long long hwidHash;
    unsigned long long hwidExHash;
    const char *authToken;
    unsigned char isDebug;
    const char *branch;
    unsigned int build;
    const char *cdnUrl;
    unsigned long long passwordHash;
    const char *ip;
    long long discordUserID;
    const char *cloudAuthHash;
} ConnectionInfo;

typedef struct weapon {
    unsigned int hash;
    unsigned char tintIndex;
    Array components;
} Weapon;

typedef struct headBlendData {
    unsigned int shapeFirstID;
    unsigned int shapeSecondID;
    unsigned int shapeThirdID;
    unsigned int skinFirstID;
    unsigned int skinSecondID;
    unsigned int skinThirdID;
    float shapeMix;
    float skinMix;
    float thirdMix;
} HeadBlendData;

typedef struct headOverlay {
    unsigned char index;
    float opacity;
    unsigned char colorType;
    unsigned char colorIndex;
    unsigned char secondColorIndex;
} HeadOverlay;

typedef struct vector2 {
    float x;
    float y;
} Vector2;

typedef struct vehicleModelInfo {
   const char *title;
   unsigned char modelType;
   unsigned char wheelsCount;
   // bool
   unsigned char hasArmoredWindows;
   unsigned char hasAutoAttachTrailer;
   Array bones;

   unsigned char primaryColor;
   unsigned char secondaryColor;
   unsigned char pearlColor;
   unsigned char wheelsColor;
   unsigned char interiorColor;
   unsigned char dashboardColor;

   unsigned short modKits[2];
   unsigned short extras;
   unsigned short defaultExtras;
} VehicleModelInfo;

typedef struct pedModelInfo {
    unsigned int hash;
    const char* name;
    Array bones;
} PedModelInfo;

typedef struct boneInfo {
    unsigned short id;
    unsigned short index;
    const char* name;
} BoneInfo;

#include "SDK.h"