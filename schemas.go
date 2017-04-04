package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. The 'id' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'id'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "id": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "temperature": {
                                "description": "Temperature timestamp.",
                                "type": "number"
                            },
                            "zone": {
                                "description": "Zone timestamp.",
                                "type": "number"
                            }
                        },
                        "required": [
                            "id"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAsset function",
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'id'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an 'id' for use as an argument to read or delete.",
                        "properties": {
                            "id": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "status": {
                                "default": 0,
                                "description": "The status of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "version"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. The arg is an 'id' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an 'id' for use as an argument to read or delete.",
                        "properties": {
                            "id": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                    "properties": {
                        "id": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "temperature": {
                            "description": "Temperature timestamp.",
                            "type": "number"
                        },
                        "zone": {
                            "description": "Zone timestamp.",
                            "type": "number"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. The 'id' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'id'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "id": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "temperature": {
                                "description": "Temperature timestamp.",
                                "type": "number"
                            },
                            "zone": {
                                "description": "Zone timestamp.",
                                "type": "number"
                            }
                        },
                        "required": [
                            "id"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "event": {
            "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'id'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
            "properties": {
                "id": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "temperature": {
                    "description": "Temperature timestamp.",
                    "type": "number"
                },
                "zone": {
                    "description": "Zone timestamp.",
                    "type": "number"
                }
            },
            "required": [
                "id"
            ],
            "type": "object"
        },
        "idKey": {
            "description": "An object containing only an 'id' for use as an argument to read or delete.",
            "properties": {
                "id": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "idandCount": {
            "description": "Requested 'id' with item 'count'.",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "id": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "id"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "status": {
                    "default": 0,
                    "description": "The status of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
            "properties": {
                "id": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "temperature": {
                    "description": "Temperature timestamp.",
                    "type": "number"
                },
                "zone": {
                    "description": "Zone timestamp.",
                    "type": "number"
                }
            },
            "type": "object"
        }
    }
}`