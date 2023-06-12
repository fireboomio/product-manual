export interface Admin__AppVersion__CreateOneAppInput {
    type:freetalk_AppType,
    version:string,
    isForce:boolean,
    latest:boolean,
}
export interface Admin__AppVersion__CreateOneAppInternalInput {
    version:string,
    createdAt:string,
    isForce:boolean,
    latest:boolean,
    type:freetalk_AppType,
    updatedAt:string,
}
export interface Admin__AppVersion__CreateOneAppResponseData {
    data?:Admin__AppVersion__CreateOneAppResponseData_data,
}
export interface Admin__AppVersion__CreateOneAppResponseData_data {
    id:string,
}
export interface Admin__AppVersion__DeleteManyAppInput {
    ids:string[],
}
export interface Admin__AppVersion__DeleteManyAppInternalInput {
    ids:string[],
}
export interface Admin__AppVersion__DeleteManyAppResponseData {
    data?:Admin__AppVersion__DeleteManyAppResponseData_data,
}
export interface Admin__AppVersion__DeleteManyAppResponseData_data {
    count:number,
}
export interface Admin__AppVersion__UpdateOneAppInput {
    description?:string,
    downloadUrl?:string,
    id:string,
    isForce?:boolean,
    latest?:boolean,
    type?:freetalk_AppType,
    version?:string,
}
export interface Admin__AppVersion__UpdateOneAppInternalInput {
    version?:string,
    description?:string,
    downloadUrl?:string,
    id:string,
    isForce?:boolean,
    latest?:boolean,
    type?:freetalk_AppType,
    updateTime:string,
}
export interface Admin__AppVersion__UpdateOneAppResponseData {
    data?:Admin__AppVersion__UpdateOneAppResponseData_data,
}
export interface Admin__AppVersion__UpdateOneAppResponseData_data {
    id:string,
}

export interface freetalk_BoolFilter {
    equals?:boolean,
    not?:freetalk_NestedBoolFilter,
}

export interface freetalk_DateTimeFilter {
    gt?:string,
    gte?:string,
    in?:string[],
    lt?:string,
    lte?:string,
    not?:freetalk_NestedDateTimeFilter,
    notIn?:string[],
    equals?:string,
}
export interface freetalk_DateTimeNullableFilter {
    equals?:string,
    gt?:string,
    gte?:string,
    in?:string[],
    lt?:string,
    lte?:string,
    not?:freetalk_NestedDateTimeNullableFilter,
    notIn?:string[],
}
export interface freetalk_IntNullableFilter {
    equals?:number,
    gt?:number,
    gte?:number,
    in?:number[],
    lt?:number,
    lte?:number,
    not?:freetalk_NestedIntNullableFilter,
    notIn?:number[],
}
export interface freetalk_NestedBoolFilter {
    equals?:boolean,
    not?:freetalk_NestedBoolFilter,
}
export interface freetalk_NestedDateTimeFilter {
    not?:freetalk_NestedDateTimeFilter,
    notIn?:string[],
    equals?:string,
    gt?:string,
    gte?:string,
    in?:string[],
    lt?:string,
    lte?:string,
}
export interface freetalk_NestedDateTimeNullableFilter {
    lt?:string,
    lte?:string,
    not?:freetalk_NestedDateTimeNullableFilter,
    notIn?:string[],
    equals?:string,
    gt?:string,
    gte?:string,
    in?:string[],
}
export interface freetalk_NestedIntNullableFilter {
    equals?:number,
    gt?:number,
    gte?:number,
    in?:number[],
    lt?:number,
    lte?:number,
    not?:freetalk_NestedIntNullableFilter,
    notIn?:number[],
}
export interface freetalk_NestedStringFilter {
    startsWith?:string,
    equals?:string,
    gt?:string,
    lt?:string,
    lte?:string,
    notIn?:string[],
    contains?:string,
    endsWith?:string,
    gte?:string,
    in?:string[],
    not?:freetalk_NestedStringFilter,
}
export interface freetalk_NestedStringNullableFilter {
    not?:freetalk_NestedStringNullableFilter,
    contains?:string,
    gte?:string,
    gt?:string,
    in?:string[],
    lt?:string,
    lte?:string,
    notIn?:string[],
    startsWith?:string,
    endsWith?:string,
    equals?:string,
}
export interface freetalk_NestedUuidFilter {
    not?:freetalk_NestedUuidFilter,
    notIn?:string[],
    equals?:string,
    gt?:string,
    gte?:string,
    in?:string[],
    lt?:string,
    lte?:string,
}
export interface freetalk_StringFilter {
    endsWith?:string,
    mode?:freetalk_QueryMode,
    not?:freetalk_NestedStringFilter,
    startsWith?:string,
    contains?:string,
    equals?:string,
    gt?:string,
    gte?:string,
    in?:string[],
    lt?:string,
    lte?:string,
    notIn?:string[],
}
export interface freetalk_StringNullableFilter {
    contains?:string,
    equals?:string,
    gte?:string,
    lt?:string,
    endsWith?:string,
    gt?:string,
    in?:string[],
    lte?:string,
    mode?:freetalk_QueryMode,
    not?:freetalk_NestedStringNullableFilter,
    notIn?:string[],
    startsWith?:string,
}
export interface freetalk_UuidFilter {
    equals?:string,
    gt?:string,
    in?:string[],
    not?:freetalk_NestedUuidFilter,
    gte?:string,
    lt?:string,
    lte?:string,
    mode?:freetalk_QueryMode,
    notIn?:string[],
}

enum freetalk_AppType {
    'Android',
    'IOS',
}

enum freetalk_QueryMode {
    'default',
    'insensitive',
}
