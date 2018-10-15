package dictionary

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResponse_GetDefinitions(t *testing.T) {
	t.Run("1 definition in array", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "metadata": {
        "provider": "Oxford University Press"
    },
    "results": [
        {
            "id": "line",
            "language": "en",
            "lexicalEntries": [
                {
                    "entries": [
                        {
                            "etymologies": [
                                "Old English līne ‘rope, series’, probably of Germanic origin, from Latin linea (fibra) ‘flax (fibre)’, from Latin linum ‘flax’, reinforced in Middle English by Old French ligne, based on Latin linea"
                            ],
                            "grammaticalFeatures": [
                                {
                                    "text": "Singular",
                                    "type": "Number"
                                }
                            ],
                            "homographNumber": "100",
                            "senses": [
                                {
                                    "definitions": [
                                        "a long, narrow mark or band"
                                    ],
                                    "examples": [
                                        {
                                            "text": "I can't draw a straight line"
                                        },
                                        {
                                            "text": "a row of closely spaced dots will look like a continuous line"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.008",
                                    "short_definitions": [
                                        "long mark or band"
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "a straight or curved continuous extent of length without breadth."
                                            ],
                                            "domains": [
                                                "Mathematics"
                                            ],
                                            "id": "m_en_gbus0583770.011",
                                            "short_definitions": [
                                                "continuous extent of length without breadth"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a direct course"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "the ball rose in a straight line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.012",
                                            "short_definitions": [
                                                "direct course"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.010"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a furrow or wrinkle in the skin, especially on the face"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "there were new lines round her eyes and mouth"
                                                },
                                                {
                                                    "text": "laughter lines"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.013",
                                            "short_definitions": [
                                                "wrinkle"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.002"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a contour or outline considered as a feature of design or composition"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "crisp architectural lines"
                                                },
                                                {
                                                    "text": "the artist's use of clean line and colour"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.014",
                                            "short_definitions": [
                                                "contour or outline"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.003"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "(on a map or graph) a curve connecting all points having a specified config property."
                                            ],
                                            "domains": [
                                                "Geography",
                                                "Mathematics"
                                            ],
                                            "id": "m_en_gbus0583770.015",
                                            "short_definitions": [
                                                "curve connecting points on map or graph"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a line marking the starting or finishing point in a race"
                                            ],
                                            "domains": [
                                                "Sport"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "a good position at the start line will put you in the front rank on the first leg"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.016",
                                            "short_definitions": [
                                                "starting or finishing point in race"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "(in football, hockey, etc.) the goal line"
                                            ],
                                            "domains": [
                                                "Sport"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "video evidence suggests the ball did not cross the line"
                                                },
                                                {
                                                    "text": "Dunne was on hand to bundle the ball over the line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.017",
                                            "short_definitions": [
                                                "(in football, hockey, etc.) goal line"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "the equator."
                                            ],
                                            "domains": [
                                                "Geography"
                                            ],
                                            "id": "m_en_gbus0583770.019",
                                            "notes": [
                                                {
                                                    "text": "\"the Line\"",
                                                    "type": "wordFormNote"
                                                }
                                            ],
                                            "short_definitions": [
                                                "equator"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a notional limit or boundary"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "television blurs the line between news and entertainment"
                                                },
                                                {
                                                    "text": "the issue of peace cut across class lines"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.022",
                                            "short_definitions": [
                                                "notional limit or boundary"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.004"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "each of the very narrow horizontal sections forming a television picture."
                                            ],
                                            "domains": [
                                                "Broadcasting"
                                            ],
                                            "id": "m_en_gbus0583770.023",
                                            "short_definitions": [
                                                "each of narrow horizontal sections forming television picture"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a narrow range of the spectrum that is noticeably brighter or darker than the adjacent parts."
                                            ],
                                            "domains": [
                                                "Physics"
                                            ],
                                            "id": "m_en_gbus0583770.024",
                                            "short_definitions": [
                                                "narrow range of spectrum"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "the level of the base of most letters, such as h and x, in printing and writing."
                                            ],
                                            "domains": [
                                                "Printing"
                                            ],
                                            "id": "m_en_gbus0583770.025",
                                            "notes": [
                                                {
                                                    "text": "\"the line\"",
                                                    "type": "wordFormNote"
                                                }
                                            ],
                                            "short_definitions": [
                                                "level of base of letters"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "denoting an image consisting of lines and solid areas, with no gradation of tone"
                                            ],
                                            "domains": [
                                                "Computing",
                                                "Printing"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "a line block"
                                                },
                                                {
                                                    "text": "line art"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.027",
                                            "notes": [
                                                {
                                                    "text": "as modifier",
                                                    "type": "grammaticalNote"
                                                }
                                            ],
                                            "short_definitions": [
                                                "referring to image consisting only of lines and solid areas"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "each of (usually five) horizontal lines forming a stave in musical notation."
                                            ],
                                            "domains": [
                                                "Music"
                                            ],
                                            "id": "m_en_gbus0583770.029",
                                            "short_definitions": [
                                                "each of horizontal lines forming musical stave"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a sequence of notes or tones forming an instrumental or vocal melody"
                                            ],
                                            "domains": [
                                                "Music"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "a powerful melodic line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.030",
                                            "short_definitions": [
                                                "sequence of notes"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a dose of a powdered narcotic drug, especially cocaine, laid out in a line ready to be taken."
                                            ],
                                            "domains": [
                                                "Narcotics"
                                            ],
                                            "id": "m_en_gbus0583770.031",
                                            "registers": [
                                                "informal"
                                            ],
                                            "short_definitions": [
                                                "dose of cocaine"
                                            ]
                                        }
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008794.001"
                                        }
                                    ]
                                },
                                {
                                    "definitions": [
                                        "a length of cord, rope, wire, or other material serving a particular purpose"
                                    ],
                                    "domains": [
                                        "Electrical"
                                    ],
                                    "examples": [
                                        {
                                            "text": "Lily pegged the washing on the line"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.033",
                                    "short_definitions": [
                                        "length of cord, rope, etc."
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "a length of sterile tubing inserted into a vein or artery in order to provide temporary access, typically so as to administer fluids or withdraw blood"
                                            ],
                                            "domains": [
                                                "Medicine"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "he's having an op this morning to put a line in his chest for IV drugs"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.209",
                                            "short_definitions": [
                                                "length of sterile tubing inserted into vein or artery"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a telephone connection or service"
                                            ],
                                            "domains": [
                                                "Telecommunications"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "I've got Inspector Jackson on the line for you"
                                                },
                                                {
                                                    "text": "a freephone advice line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.035",
                                            "short_definitions": [
                                                "telephone connection"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a railway track"
                                            ],
                                            "domains": [
                                                "Railways"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "passengers were hit by delays caused by leaves on the line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.036",
                                            "short_definitions": [
                                                "railway track"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a branch or route of a railway system"
                                            ],
                                            "domains": [
                                                "Railways"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "the Glasgow to London line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.037",
                                            "short_definitions": [
                                                "route of railway system"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a company that provides ships, aircraft, or buses on particular routes on a regular basis"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "a major shipping line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.038",
                                            "short_definitions": [
                                                "transport company"
                                            ]
                                        }
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008794.006"
                                        }
                                    ]
                                },
                                {
                                    "definitions": [
                                        "a horizontal row of written or printed words"
                                    ],
                                    "examples": [
                                        {
                                            "text": "take the cursor up one line and press the delete key"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.040",
                                    "short_definitions": [
                                        "horizontal row of words"
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "a part of a poem or song forming one row of written or printed words"
                                            ],
                                            "domains": [
                                                "Prosody"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "each stanza has eight lines"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.041",
                                            "short_definitions": [
                                                "part of poem forming row"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.018"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "the words of an actor's part in a play or film"
                                            ],
                                            "domains": [
                                                "Theatre"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "he couldn't seem to remember his lines and had to read his dialogue off boards"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.042",
                                            "notes": [
                                                {
                                                    "text": "\"lines\"",
                                                    "type": "wordFormNote"
                                                }
                                            ],
                                            "short_definitions": [
                                                "actor's words"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.014"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "an amount of text or number of repetitions of a sentence written out as a school punishment"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "five hundred lines to anyone caught sneaking in before the bell!"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.045",
                                            "notes": [
                                                {
                                                    "text": "\"lines\"",
                                                    "type": "wordFormNote"
                                                }
                                            ],
                                            "regions": [
                                                "British"
                                            ],
                                            "short_definitions": [
                                                "amount of text copied out as school punishment"
                                            ]
                                        }
                                    ]
                                },
                                {
                                    "definitions": [
                                        "a row of people or things"
                                    ],
                                    "examples": [
                                        {
                                            "text": "a line of altar boys proceeded down the aisle"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.047",
                                    "short_definitions": [
                                        "row of people or things"
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "a queue."
                                            ],
                                            "id": "m_en_gbus0583770.048",
                                            "regions": [
                                                "North American"
                                            ],
                                            "short_definitions": [
                                                "queue"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a connected series of people following one another in time (used especially of several generations of a family)"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "we follow the history of a family through the male line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.049",
                                            "short_definitions": [
                                                "connected series of people or generations"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.017"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a series of related things"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "the bill is the latest in a long line of measures to protect society from criminals"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.052",
                                            "short_definitions": [
                                                "series of related things"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.009"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a range of commercial goods"
                                            ],
                                            "domains": [
                                                "Commerce"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "the company intends to hire more people and expand its product line"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.053",
                                            "short_definitions": [
                                                "range of commercial goods"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.016"
                                                }
                                            ]
                                        }
                                    ]
                                },
                                {
                                    "definitions": [
                                        "an area or branch of activity"
                                    ],
                                    "examples": [
                                        {
                                            "text": "the stresses unique to their line of work"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.056",
                                    "short_definitions": [
                                        "sphere of activity"
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "a direction, course, or channel"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "he opened another line of attack"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.057",
                                            "short_definitions": [
                                                "direction, course, or channel"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.011"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a manner of doing or thinking about something"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "you can't run a business on these lines"
                                                },
                                                {
                                                    "text": "the superintendent was thinking along the same lines"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.058",
                                            "notes": [
                                                {
                                                    "text": "\"lines\"",
                                                    "type": "wordFormNote"
                                                }
                                            ],
                                            "short_definitions": [
                                                "manner of doing something"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.011"
                                                },
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.012"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "an agreed approach; a policy"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "the official line is that there were no chemical attacks on allied troops"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.059",
                                            "short_definitions": [
                                                "policy"
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "a false or exaggerated remark or story"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "none of my chat-up lines ever worked"
                                                },
                                                {
                                                    "text": "he fed me a line about some nightclubbing Japanese photographer"
                                                }
                                            ],
                                            "id": "m_en_gbus0583770.060",
                                            "registers": [
                                                "informal"
                                            ],
                                            "short_definitions": [
                                                "false or exaggerated story"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.013"
                                                }
                                            ]
                                        }
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008794.015"
                                        }
                                    ]
                                },
                                {
                                    "definitions": [
                                        "a connected series of military fieldworks or defences facing an enemy force"
                                    ],
                                    "domains": [
                                        "Military"
                                    ],
                                    "examples": [
                                        {
                                            "text": "raids behind enemy lines"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.062",
                                    "short_definitions": [
                                        "connected series of military fieldworks"
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "an arrangement of soldiers or ships in a column or line formation; a line of battle."
                                            ],
                                            "domains": [
                                                "Military",
                                                "Naval"
                                            ],
                                            "id": "m_en_gbus0583770.063",
                                            "short_definitions": [
                                                "arrangement of soldiers or ships in column"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008794.007"
                                                }
                                            ]
                                        },
                                        {
                                            "definitions": [
                                                "regular army regiments (as opposed to auxiliary forces or household troops)."
                                            ],
                                            "domains": [
                                                "Military"
                                            ],
                                            "id": "m_en_gbus0583770.064",
                                            "notes": [
                                                {
                                                    "text": "\"the line\"",
                                                    "type": "wordFormNote"
                                                }
                                            ],
                                            "short_definitions": [
                                                "regular army regiments"
                                            ]
                                        }
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008794.005"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Noun",
                    "pronunciations": [
                        {
                            "audioFile": "http://audio.oxforddictionaries.com/en/mp3/line_gb_1.mp3",
                            "dialects": [
                                "British English"
                            ],
                            "phoneticNotation": "IPA",
                            "phoneticSpelling": "lʌɪn"
                        }
                    ],
                    "text": "line"
                },
                {
                    "entries": [
                        {
                            "grammaticalFeatures": [
                                {
                                    "text": "Transitive",
                                    "type": "Subcategorization"
                                },
                                {
                                    "text": "Present",
                                    "type": "Tense"
                                }
                            ],
                            "homographNumber": "101",
                            "senses": [
                                {
                                    "definitions": [
                                        "stand or be positioned at intervals along"
                                    ],
                                    "examples": [
                                        {
                                            "text": "a processional route lined by people waving flags"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.067",
                                    "short_definitions": [
                                        "stand at intervals along"
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008794.030"
                                        }
                                    ]
                                },
                                {
                                    "definitions": [
                                        "mark or cover with lines"
                                    ],
                                    "examples": [
                                        {
                                            "text": "a thin woman with a lined face"
                                        },
                                        {
                                            "text": "lined paper"
                                        }
                                    ],
                                    "id": "m_en_gbus0583770.074",
                                    "notes": [
                                        {
                                            "text": "usually as adjective \"lined\"",
                                            "type": "wordFormNote"
                                        }
                                    ],
                                    "short_definitions": [
                                        "mark with lines"
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "lined",
                                            "sense_id": "t_en_gb0008798.001"
                                        },
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008794.029"
                                        },
                                        {
                                            "entry_id": "lined",
                                            "sense_id": "t_en_gb0008798.002"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "etymologies": [
                                "late Middle English: from obsolete line ‘flax’, with reference to the config use of linen for linings"
                            ],
                            "grammaticalFeatures": [
                                {
                                    "text": "Transitive",
                                    "type": "Subcategorization"
                                },
                                {
                                    "text": "Present",
                                    "type": "Tense"
                                }
                            ],
                            "homographNumber": "200",
                            "senses": [
                                {
                                    "definitions": [
                                        "cover the inside surface of (a container or garment) with a layer of different material"
                                    ],
                                    "examples": [
                                        {
                                            "text": "a basket lined with polythene"
                                        }
                                    ],
                                    "id": "m_en_gbus0583780.005",
                                    "short_definitions": [
                                        "cover inside surface of container or garment with layer of different material"
                                    ],
                                    "subsenses": [
                                        {
                                            "definitions": [
                                                "form a layer on the inside surface of (an area); cover as if with a lining"
                                            ],
                                            "examples": [
                                                {
                                                    "text": "hundreds of telegrams lined the walls"
                                                }
                                            ],
                                            "id": "m_en_gbus0583780.011",
                                            "short_definitions": [
                                                "form layer on inside surface of"
                                            ],
                                            "thesaurusLinks": [
                                                {
                                                    "entry_id": "line",
                                                    "sense_id": "t_en_gb0008795.001"
                                                }
                                            ]
                                        }
                                    ],
                                    "thesaurusLinks": [
                                        {
                                            "entry_id": "lined",
                                            "sense_id": "t_en_gb0008799.001"
                                        },
                                        {
                                            "entry_id": "line",
                                            "sense_id": "t_en_gb0008795.001"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Verb",
                    "pronunciations": [
                        {
                            "audioFile": "http://audio.oxforddictionaries.com/en/mp3/line_gb_1.mp3",
                            "dialects": [
                                "British English"
                            ],
                            "phoneticNotation": "IPA",
                            "phoneticSpelling": "lʌɪn"
                        }
                    ],
                    "text": "line"
                }
            ],
            "type": "headword",
            "word": "line"
        }
    ]
}`)
		expected := []string{"a long, narrow mark or band"}
		var response OxfordResponse
		json.Unmarshal(body, &response)
		actual := response.GetDefinitions()
		assertStringArray(t, expected, actual)
	})
}

func TestResponse_GetSynonyms(t *testing.T) {
	t.Run("More than 5 in first subSense", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "metadata": {
        "provider": "Oxford University Press"
    },
    "results": [
        {
            "id": "line",
            "language": "en",
            "lexicalEntries": [
                {
                    "entries": [
                        {
                            "homographNumber": "100",
                            "senses": [
                                {
                                    "examples": [
                                        {
                                            "text": "he drew a line through the name"
                                        },
                                        {
                                            "text": "a pattern of wavy lines"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.001",
                                    "subsenses": [
                                        {
                                            "id": "id498a39e3-d8be-49a3-a8cd-1ccd55240845",
                                            "synonyms": [
                                                {
                                                    "id": "underline",
                                                    "language": "en",
                                                    "text": "underline"
                                                },
                                                {
                                                    "id": "underscore",
                                                    "language": "en",
                                                    "text": "underscore"
                                                },
                                                {
                                                    "id": "stroke",
                                                    "language": "en",
                                                    "text": "stroke"
                                                },
                                                {
                                                    "id": "slash",
                                                    "language": "en",
                                                    "text": "slash"
                                                },
                                                {
                                                    "id": "virgule",
                                                    "language": "en",
                                                    "text": "virgule"
                                                },
                                                {
                                                    "id": "solidus",
                                                    "language": "en",
                                                    "text": "solidus"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id4cef684b-4862-4593-b916-1052e83d914f",
                                            "synonyms": [
                                                {
                                                    "id": "stripe",
                                                    "language": "en",
                                                    "text": "stripe"
                                                },
                                                {
                                                    "id": "strip",
                                                    "language": "en",
                                                    "text": "strip"
                                                },
                                                {
                                                    "id": "band",
                                                    "language": "en",
                                                    "text": "band"
                                                },
                                                {
                                                    "id": "streak",
                                                    "language": "en",
                                                    "text": "streak"
                                                },
                                                {
                                                    "id": "belt",
                                                    "language": "en",
                                                    "text": "belt"
                                                },
                                                {
                                                    "id": "striation",
                                                    "language": "en",
                                                    "text": "striation"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id111109d6-8c7f-4c9d-8bd4-a7289b3b6fc8",
                                            "registers": [
                                                "technical"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "stria",
                                                    "language": "en",
                                                    "text": "stria"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id34ced156-f5cd-4805-84f7-04eece21808c",
                                            "regions": [
                                                "British"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "oblique",
                                                    "language": "en",
                                                    "text": "oblique"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "dash",
                                            "language": "en",
                                            "text": "dash"
                                        },
                                        {
                                            "id": "rule",
                                            "language": "en",
                                            "text": "rule"
                                        },
                                        {
                                            "id": "bar",
                                            "language": "en",
                                            "text": "bar"
                                        },
                                        {
                                            "id": "score",
                                            "language": "en",
                                            "text": "score"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "there were new lines round her eyes and mouth"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.002",
                                    "subsenses": [
                                        {
                                            "id": "id91c8dc1b-ba78-42e1-ba37-278e00d5f53f",
                                            "synonyms": [
                                                {
                                                    "id": "scar",
                                                    "language": "en",
                                                    "text": "scar"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "wrinkle",
                                            "language": "en",
                                            "text": "wrinkle"
                                        },
                                        {
                                            "id": "furrow",
                                            "language": "en",
                                            "text": "furrow"
                                        },
                                        {
                                            "id": "crease",
                                            "language": "en",
                                            "text": "crease"
                                        },
                                        {
                                            "id": "crinkle",
                                            "language": "en",
                                            "text": "crinkle"
                                        },
                                        {
                                            "id": "crow%27s_foot",
                                            "language": "en",
                                            "text": "crow's foot"
                                        },
                                        {
                                            "id": "groove",
                                            "language": "en",
                                            "text": "groove"
                                        },
                                        {
                                            "id": "corrugation",
                                            "language": "en",
                                            "text": "corrugation"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "the classic lines of its exterior"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.003",
                                    "synonyms": [
                                        {
                                            "id": "contour",
                                            "language": "en",
                                            "text": "contour"
                                        },
                                        {
                                            "id": "outline",
                                            "language": "en",
                                            "text": "outline"
                                        },
                                        {
                                            "id": "configuration",
                                            "language": "en",
                                            "text": "configuration"
                                        },
                                        {
                                            "id": "shape",
                                            "language": "en",
                                            "text": "shape"
                                        },
                                        {
                                            "id": "figure",
                                            "language": "en",
                                            "text": "figure"
                                        },
                                        {
                                            "id": "delineation",
                                            "language": "en",
                                            "text": "delineation"
                                        },
                                        {
                                            "id": "silhouette",
                                            "language": "en",
                                            "text": "silhouette"
                                        },
                                        {
                                            "id": "profile",
                                            "language": "en",
                                            "text": "profile"
                                        },
                                        {
                                            "id": "features",
                                            "language": "en",
                                            "text": "features"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "he headed the ball over the line"
                                        },
                                        {
                                            "text": "the county line"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.004",
                                    "synonyms": [
                                        {
                                            "id": "boundary",
                                            "language": "en",
                                            "text": "boundary"
                                        },
                                        {
                                            "id": "boundary_line",
                                            "language": "en",
                                            "text": "boundary line"
                                        },
                                        {
                                            "id": "limit",
                                            "language": "en",
                                            "text": "limit"
                                        },
                                        {
                                            "id": "border",
                                            "language": "en",
                                            "text": "border"
                                        },
                                        {
                                            "id": "borderline",
                                            "language": "en",
                                            "text": "borderline"
                                        },
                                        {
                                            "id": "bound",
                                            "language": "en",
                                            "text": "bound"
                                        },
                                        {
                                            "id": "bounding_line",
                                            "language": "en",
                                            "text": "bounding line"
                                        },
                                        {
                                            "id": "frontier",
                                            "language": "en",
                                            "text": "frontier"
                                        },
                                        {
                                            "id": "partition",
                                            "language": "en",
                                            "text": "partition"
                                        },
                                        {
                                            "id": "demarcation_line",
                                            "language": "en",
                                            "text": "demarcation line"
                                        },
                                        {
                                            "id": "dividing_line",
                                            "language": "en",
                                            "text": "dividing line"
                                        },
                                        {
                                            "id": "end_point",
                                            "language": "en",
                                            "text": "end point"
                                        },
                                        {
                                            "id": "cut-off_point",
                                            "language": "en",
                                            "text": "cut-off point"
                                        },
                                        {
                                            "id": "termination",
                                            "language": "en",
                                            "text": "termination"
                                        },
                                        {
                                            "id": "edge",
                                            "language": "en",
                                            "text": "edge"
                                        },
                                        {
                                            "id": "pale",
                                            "language": "en",
                                            "text": "pale"
                                        },
                                        {
                                            "id": "margin",
                                            "language": "en",
                                            "text": "margin"
                                        },
                                        {
                                            "id": "perimeter",
                                            "language": "en",
                                            "text": "perimeter"
                                        },
                                        {
                                            "id": "periphery",
                                            "language": "en",
                                            "text": "periphery"
                                        },
                                        {
                                            "id": "rim",
                                            "language": "en",
                                            "text": "rim"
                                        },
                                        {
                                            "id": "extremity",
                                            "language": "en",
                                            "text": "extremity"
                                        },
                                        {
                                            "id": "fringe",
                                            "language": "en",
                                            "text": "fringe"
                                        },
                                        {
                                            "id": "threshold",
                                            "language": "en",
                                            "text": "threshold"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "they were behind enemy lines"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.005",
                                    "subsenses": [
                                        {
                                            "id": "id033cf20d-6d83-4602-9fb3-32f1571d4f08",
                                            "synonyms": [
                                                {
                                                    "id": "trenches",
                                                    "language": "en",
                                                    "text": "trenches"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "position",
                                            "language": "en",
                                            "text": "position"
                                        },
                                        {
                                            "id": "formation",
                                            "language": "en",
                                            "text": "formation"
                                        },
                                        {
                                            "id": "disposition",
                                            "language": "en",
                                            "text": "disposition"
                                        },
                                        {
                                            "id": "front",
                                            "language": "en",
                                            "text": "front"
                                        },
                                        {
                                            "id": "front_line",
                                            "language": "en",
                                            "text": "front line"
                                        },
                                        {
                                            "id": "firing_line",
                                            "language": "en",
                                            "text": "firing line"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "a fishing line"
                                        },
                                        {
                                            "text": "he put the washing on the line"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.006",
                                    "synonyms": [
                                        {
                                            "id": "cord",
                                            "language": "en",
                                            "text": "cord"
                                        },
                                        {
                                            "id": "rope",
                                            "language": "en",
                                            "text": "rope"
                                        },
                                        {
                                            "id": "string",
                                            "language": "en",
                                            "text": "string"
                                        },
                                        {
                                            "id": "cable",
                                            "language": "en",
                                            "text": "cable"
                                        },
                                        {
                                            "id": "wire",
                                            "language": "en",
                                            "text": "wire"
                                        },
                                        {
                                            "id": "thread",
                                            "language": "en",
                                            "text": "thread"
                                        },
                                        {
                                            "id": "twine",
                                            "language": "en",
                                            "text": "twine"
                                        },
                                        {
                                            "id": "strand",
                                            "language": "en",
                                            "text": "strand"
                                        },
                                        {
                                            "id": "filament",
                                            "language": "en",
                                            "text": "filament"
                                        },
                                        {
                                            "id": "ligature",
                                            "language": "en",
                                            "text": "ligature"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "a line of soldiers"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.007",
                                    "subsenses": [
                                        {
                                            "id": "idcadd6955-b626-4ae0-851a-9b332a489faf",
                                            "synonyms": [
                                                {
                                                    "id": "train",
                                                    "language": "en",
                                                    "text": "train"
                                                },
                                                {
                                                    "id": "convoy",
                                                    "language": "en",
                                                    "text": "convoy"
                                                },
                                                {
                                                    "id": "procession",
                                                    "language": "en",
                                                    "text": "procession"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "ida1ef90e8-b5cf-4f04-bc4e-3863e849c53e",
                                            "synonyms": [
                                                {
                                                    "id": "row",
                                                    "language": "en",
                                                    "text": "row"
                                                },
                                                {
                                                    "id": "queue",
                                                    "language": "en",
                                                    "text": "queue"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id62382bde-24f9-4ea4-bb0a-cb5943a0582e",
                                            "regions": [
                                                "British"
                                            ],
                                            "registers": [
                                                "informal"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "crocodile",
                                                    "language": "en",
                                                    "text": "crocodile"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "file",
                                            "language": "en",
                                            "text": "file"
                                        },
                                        {
                                            "id": "rank",
                                            "language": "en",
                                            "text": "rank"
                                        },
                                        {
                                            "id": "column",
                                            "language": "en",
                                            "text": "column"
                                        },
                                        {
                                            "id": "string",
                                            "language": "en",
                                            "text": "string"
                                        },
                                        {
                                            "id": "chain",
                                            "language": "en",
                                            "text": "chain"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "a line of figures"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.008",
                                    "synonyms": [
                                        {
                                            "id": "column",
                                            "language": "en",
                                            "text": "column"
                                        },
                                        {
                                            "id": "row",
                                            "language": "en",
                                            "text": "row"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "it seemed to be the latest in a long line of crass decisions"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.009",
                                    "subsenses": [
                                        {
                                            "id": "id17047fb4-35ca-4c81-9fc9-c58abd02731e",
                                            "synonyms": [
                                                {
                                                    "id": "progression",
                                                    "language": "en",
                                                    "text": "progression"
                                                },
                                                {
                                                    "id": "course",
                                                    "language": "en",
                                                    "text": "course"
                                                },
                                                {
                                                    "id": "set",
                                                    "language": "en",
                                                    "text": "set"
                                                },
                                                {
                                                    "id": "cycle",
                                                    "language": "en",
                                                    "text": "cycle"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "series",
                                            "language": "en",
                                            "text": "series"
                                        },
                                        {
                                            "id": "sequence",
                                            "language": "en",
                                            "text": "sequence"
                                        },
                                        {
                                            "id": "succession",
                                            "language": "en",
                                            "text": "succession"
                                        },
                                        {
                                            "id": "chain",
                                            "language": "en",
                                            "text": "chain"
                                        },
                                        {
                                            "id": "string",
                                            "language": "en",
                                            "text": "string"
                                        },
                                        {
                                            "id": "train",
                                            "language": "en",
                                            "text": "train"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "it stopped right in the line of flight of some bees"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.010",
                                    "subsenses": [
                                        {
                                            "id": "idd99b7933-9cdb-46c8-b725-ceecfdff65d9",
                                            "synonyms": [
                                                {
                                                    "id": "trajectory",
                                                    "language": "en",
                                                    "text": "trajectory"
                                                },
                                                {
                                                    "id": "bearing",
                                                    "language": "en",
                                                    "text": "bearing"
                                                },
                                                {
                                                    "id": "orientation",
                                                    "language": "en",
                                                    "text": "orientation"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "course",
                                            "language": "en",
                                            "text": "course"
                                        },
                                        {
                                            "id": "route",
                                            "language": "en",
                                            "text": "route"
                                        },
                                        {
                                            "id": "track",
                                            "language": "en",
                                            "text": "track"
                                        },
                                        {
                                            "id": "channel",
                                            "language": "en",
                                            "text": "channel"
                                        },
                                        {
                                            "id": "path",
                                            "language": "en",
                                            "text": "path"
                                        },
                                        {
                                            "id": "way",
                                            "language": "en",
                                            "text": "way"
                                        },
                                        {
                                            "id": "run",
                                            "language": "en",
                                            "text": "run"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "they took a very tough line with the industry right from the word go"
                                        },
                                        {
                                            "text": "ministers are obliged to follow the party line"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.011",
                                    "subsenses": [
                                        {
                                            "id": "id37cdd666-ec04-4904-be64-402412955eea",
                                            "synonyms": [
                                                {
                                                    "id": "policy",
                                                    "language": "en",
                                                    "text": "policy"
                                                },
                                                {
                                                    "id": "practice",
                                                    "language": "en",
                                                    "text": "practice"
                                                },
                                                {
                                                    "id": "scheme",
                                                    "language": "en",
                                                    "text": "scheme"
                                                },
                                                {
                                                    "id": "approach",
                                                    "language": "en",
                                                    "text": "approach"
                                                },
                                                {
                                                    "id": "plan",
                                                    "language": "en",
                                                    "text": "plan"
                                                },
                                                {
                                                    "id": "programme",
                                                    "language": "en",
                                                    "text": "programme"
                                                },
                                                {
                                                    "id": "position",
                                                    "language": "en",
                                                    "text": "position"
                                                },
                                                {
                                                    "id": "stance",
                                                    "language": "en",
                                                    "text": "stance"
                                                },
                                                {
                                                    "id": "philosophy",
                                                    "language": "en",
                                                    "text": "philosophy"
                                                },
                                                {
                                                    "id": "argument",
                                                    "language": "en",
                                                    "text": "argument"
                                                },
                                                {
                                                    "id": "avenue",
                                                    "language": "en",
                                                    "text": "avenue"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id8f2b03af-9ceb-4c7e-9d34-8a60fb8fa0f2",
                                            "synonyms": [
                                                {
                                                    "id": "modus_operandi",
                                                    "language": "en",
                                                    "text": "modus operandi"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "course_of_action",
                                            "language": "en",
                                            "text": "course of action"
                                        },
                                        {
                                            "id": "course",
                                            "language": "en",
                                            "text": "course"
                                        },
                                        {
                                            "id": "procedure",
                                            "language": "en",
                                            "text": "procedure"
                                        },
                                        {
                                            "id": "mo",
                                            "language": "en",
                                            "text": "MO"
                                        },
                                        {
                                            "id": "technique",
                                            "language": "en",
                                            "text": "technique"
                                        },
                                        {
                                            "id": "way",
                                            "language": "en",
                                            "text": "way"
                                        },
                                        {
                                            "id": "tactic",
                                            "language": "en",
                                            "text": "tactic"
                                        },
                                        {
                                            "id": "tack",
                                            "language": "en",
                                            "text": "tack"
                                        },
                                        {
                                            "id": "system",
                                            "language": "en",
                                            "text": "system"
                                        },
                                        {
                                            "id": "method",
                                            "language": "en",
                                            "text": "method"
                                        },
                                        {
                                            "id": "process",
                                            "language": "en",
                                            "text": "process"
                                        },
                                        {
                                            "id": "manner",
                                            "language": "en",
                                            "text": "manner"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "she had not been listening, but pursuing her own line of thought"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.012",
                                    "synonyms": [
                                        {
                                            "id": "course",
                                            "language": "en",
                                            "text": "course"
                                        },
                                        {
                                            "id": "direction",
                                            "language": "en",
                                            "text": "direction"
                                        },
                                        {
                                            "id": "drift",
                                            "language": "en",
                                            "text": "drift"
                                        },
                                        {
                                            "id": "tack",
                                            "language": "en",
                                            "text": "tack"
                                        },
                                        {
                                            "id": "tendency",
                                            "language": "en",
                                            "text": "tendency"
                                        },
                                        {
                                            "id": "trend",
                                            "language": "en",
                                            "text": "trend"
                                        },
                                        {
                                            "id": "bias",
                                            "language": "en",
                                            "text": "bias"
                                        },
                                        {
                                            "id": "tenor",
                                            "language": "en",
                                            "text": "tenor"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "oh, come on, don't give me that line"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.013",
                                    "subsenses": [
                                        {
                                            "id": "id6b216372-b6ed-4788-badc-e5a7df77e1d8",
                                            "registers": [
                                                "informal"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "spiel",
                                                    "language": "en",
                                                    "text": "spiel"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "patter",
                                            "language": "en",
                                            "text": "patter"
                                        },
                                        {
                                            "id": "story",
                                            "language": "en",
                                            "text": "story"
                                        },
                                        {
                                            "id": "pitch",
                                            "language": "en",
                                            "text": "pitch"
                                        },
                                        {
                                            "id": "piece_of_fiction",
                                            "language": "en",
                                            "text": "piece of fiction"
                                        },
                                        {
                                            "id": "fabrication",
                                            "language": "en",
                                            "text": "fabrication"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "he couldn't seem to remember his lines"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.014",
                                    "synonyms": [
                                        {
                                            "id": "words",
                                            "language": "en",
                                            "text": "words"
                                        },
                                        {
                                            "id": "role",
                                            "language": "en",
                                            "text": "role"
                                        },
                                        {
                                            "id": "part",
                                            "language": "en",
                                            "text": "part"
                                        },
                                        {
                                            "id": "script",
                                            "language": "en",
                                            "text": "script"
                                        },
                                        {
                                            "id": "speech",
                                            "language": "en",
                                            "text": "speech"
                                        },
                                        {
                                            "id": "dialogue",
                                            "language": "en",
                                            "text": "dialogue"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "there are no jobs nowadays in my line"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.015",
                                    "subsenses": [
                                        {
                                            "id": "ide0a764e0-9d71-4c6e-891c-e3a70af65863",
                                            "synonyms": [
                                                {
                                                    "id": "specialty",
                                                    "language": "en",
                                                    "text": "specialty"
                                                },
                                                {
                                                    "id": "forte",
                                                    "language": "en",
                                                    "text": "forte"
                                                },
                                                {
                                                    "id": "province",
                                                    "language": "en",
                                                    "text": "province"
                                                },
                                                {
                                                    "id": "department",
                                                    "language": "en",
                                                    "text": "department"
                                                },
                                                {
                                                    "id": "sphere",
                                                    "language": "en",
                                                    "text": "sphere"
                                                },
                                                {
                                                    "id": "area",
                                                    "language": "en",
                                                    "text": "area"
                                                },
                                                {
                                                    "id": "area_of_expertise",
                                                    "language": "en",
                                                    "text": "area of expertise"
                                                },
                                                {
                                                    "id": "domain",
                                                    "language": "en",
                                                    "text": "domain"
                                                },
                                                {
                                                    "id": "realm",
                                                    "language": "en",
                                                    "text": "realm"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "idcbf333b4-c0c4-4be7-8cdd-93d5c71b9845",
                                            "synonyms": [
                                                {
                                                    "id": "m%C3%A9tier",
                                                    "language": "en",
                                                    "text": "métier"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id4e4bd301-4c63-4398-9d6d-3453bd47c793",
                                            "registers": [
                                                "informal"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "line_of_country",
                                                    "language": "en",
                                                    "text": "line of country"
                                                },
                                                {
                                                    "id": "game",
                                                    "language": "en",
                                                    "text": "game"
                                                },
                                                {
                                                    "id": "thing",
                                                    "language": "en",
                                                    "text": "thing"
                                                },
                                                {
                                                    "id": "bag",
                                                    "language": "en",
                                                    "text": "bag"
                                                },
                                                {
                                                    "id": "pigeon",
                                                    "language": "en",
                                                    "text": "pigeon"
                                                },
                                                {
                                                    "id": "racket",
                                                    "language": "en",
                                                    "text": "racket"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "line_of_work",
                                            "language": "en",
                                            "text": "line of work"
                                        },
                                        {
                                            "id": "line_of_business",
                                            "language": "en",
                                            "text": "line of business"
                                        },
                                        {
                                            "id": "business",
                                            "language": "en",
                                            "text": "business"
                                        },
                                        {
                                            "id": "field",
                                            "language": "en",
                                            "text": "field"
                                        },
                                        {
                                            "id": "trade",
                                            "language": "en",
                                            "text": "trade"
                                        },
                                        {
                                            "id": "occupation",
                                            "language": "en",
                                            "text": "occupation"
                                        },
                                        {
                                            "id": "employment",
                                            "language": "en",
                                            "text": "employment"
                                        },
                                        {
                                            "id": "profession",
                                            "language": "en",
                                            "text": "profession"
                                        },
                                        {
                                            "id": "work",
                                            "language": "en",
                                            "text": "work"
                                        },
                                        {
                                            "id": "job",
                                            "language": "en",
                                            "text": "job"
                                        },
                                        {
                                            "id": "day_job",
                                            "language": "en",
                                            "text": "day job"
                                        },
                                        {
                                            "id": "calling",
                                            "language": "en",
                                            "text": "calling"
                                        },
                                        {
                                            "id": "vocation",
                                            "language": "en",
                                            "text": "vocation"
                                        },
                                        {
                                            "id": "career",
                                            "language": "en",
                                            "text": "career"
                                        },
                                        {
                                            "id": "pursuit",
                                            "language": "en",
                                            "text": "pursuit"
                                        },
                                        {
                                            "id": "activity",
                                            "language": "en",
                                            "text": "activity"
                                        },
                                        {
                                            "id": "walk_of_life",
                                            "language": "en",
                                            "text": "walk of life"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "he's introduced his own line of cologne"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.016",
                                    "synonyms": [
                                        {
                                            "id": "brand",
                                            "language": "en",
                                            "text": "brand"
                                        },
                                        {
                                            "id": "kind",
                                            "language": "en",
                                            "text": "kind"
                                        },
                                        {
                                            "id": "sort",
                                            "language": "en",
                                            "text": "sort"
                                        },
                                        {
                                            "id": "type",
                                            "language": "en",
                                            "text": "type"
                                        },
                                        {
                                            "id": "variety",
                                            "language": "en",
                                            "text": "variety"
                                        },
                                        {
                                            "id": "make",
                                            "language": "en",
                                            "text": "make"
                                        },
                                        {
                                            "id": "label",
                                            "language": "en",
                                            "text": "label"
                                        },
                                        {
                                            "id": "trade_name",
                                            "language": "en",
                                            "text": "trade name"
                                        },
                                        {
                                            "id": "trademark",
                                            "language": "en",
                                            "text": "trademark"
                                        },
                                        {
                                            "id": "registered_trademark",
                                            "language": "en",
                                            "text": "registered trademark"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "a man from a noble line claiming royal descent"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.017",
                                    "subsenses": [
                                        {
                                            "id": "id9d7a0596-75df-480f-adc2-b63aa3710344",
                                            "synonyms": [
                                                {
                                                    "id": "stock",
                                                    "language": "en",
                                                    "text": "stock"
                                                },
                                                {
                                                    "id": "strain",
                                                    "language": "en",
                                                    "text": "strain"
                                                },
                                                {
                                                    "id": "race",
                                                    "language": "en",
                                                    "text": "race"
                                                },
                                                {
                                                    "id": "bloodline",
                                                    "language": "en",
                                                    "text": "bloodline"
                                                },
                                                {
                                                    "id": "blood",
                                                    "language": "en",
                                                    "text": "blood"
                                                },
                                                {
                                                    "id": "breeding",
                                                    "language": "en",
                                                    "text": "breeding"
                                                },
                                                {
                                                    "id": "pedigree",
                                                    "language": "en",
                                                    "text": "pedigree"
                                                },
                                                {
                                                    "id": "succession",
                                                    "language": "en",
                                                    "text": "succession"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "ancestry",
                                            "language": "en",
                                            "text": "ancestry"
                                        },
                                        {
                                            "id": "family",
                                            "language": "en",
                                            "text": "family"
                                        },
                                        {
                                            "id": "parentage",
                                            "language": "en",
                                            "text": "parentage"
                                        },
                                        {
                                            "id": "birth",
                                            "language": "en",
                                            "text": "birth"
                                        },
                                        {
                                            "id": "descent",
                                            "language": "en",
                                            "text": "descent"
                                        },
                                        {
                                            "id": "lineage",
                                            "language": "en",
                                            "text": "lineage"
                                        },
                                        {
                                            "id": "extraction",
                                            "language": "en",
                                            "text": "extraction"
                                        },
                                        {
                                            "id": "derivation",
                                            "language": "en",
                                            "text": "derivation"
                                        },
                                        {
                                            "id": "heritage",
                                            "language": "en",
                                            "text": "heritage"
                                        },
                                        {
                                            "id": "genealogy",
                                            "language": "en",
                                            "text": "genealogy"
                                        },
                                        {
                                            "id": "roots",
                                            "language": "en",
                                            "text": "roots"
                                        },
                                        {
                                            "id": "house",
                                            "language": "en",
                                            "text": "house"
                                        },
                                        {
                                            "id": "dynasty",
                                            "language": "en",
                                            "text": "dynasty"
                                        },
                                        {
                                            "id": "origin",
                                            "language": "en",
                                            "text": "origin"
                                        },
                                        {
                                            "id": "background",
                                            "language": "en",
                                            "text": "background"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "the opening line of Wilfred Owen's ‘Anthem for Doomed Youth’"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.018",
                                    "subsenses": [
                                        {
                                            "id": "id3ac8710f-8bce-474c-93b2-1fcc335ad63a",
                                            "synonyms": [
                                                {
                                                    "id": "passage",
                                                    "language": "en",
                                                    "text": "passage"
                                                },
                                                {
                                                    "id": "extract",
                                                    "language": "en",
                                                    "text": "extract"
                                                },
                                                {
                                                    "id": "quotation",
                                                    "language": "en",
                                                    "text": "quotation"
                                                },
                                                {
                                                    "id": "quote",
                                                    "language": "en",
                                                    "text": "quote"
                                                },
                                                {
                                                    "id": "citation",
                                                    "language": "en",
                                                    "text": "citation"
                                                },
                                                {
                                                    "id": "section",
                                                    "language": "en",
                                                    "text": "section"
                                                },
                                                {
                                                    "id": "piece",
                                                    "language": "en",
                                                    "text": "piece"
                                                },
                                                {
                                                    "id": "part",
                                                    "language": "en",
                                                    "text": "part"
                                                },
                                                {
                                                    "id": "snippet",
                                                    "language": "en",
                                                    "text": "snippet"
                                                },
                                                {
                                                    "id": "sound_bite",
                                                    "language": "en",
                                                    "text": "sound bite"
                                                },
                                                {
                                                    "id": "fragment",
                                                    "language": "en",
                                                    "text": "fragment"
                                                },
                                                {
                                                    "id": "portion",
                                                    "language": "en",
                                                    "text": "portion"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "sentence",
                                            "language": "en",
                                            "text": "sentence"
                                        },
                                        {
                                            "id": "phrase",
                                            "language": "en",
                                            "text": "phrase"
                                        },
                                        {
                                            "id": "group_of_words",
                                            "language": "en",
                                            "text": "group of words"
                                        },
                                        {
                                            "id": "prosodic_unit",
                                            "language": "en",
                                            "text": "prosodic unit"
                                        },
                                        {
                                            "id": "construction",
                                            "language": "en",
                                            "text": "construction"
                                        },
                                        {
                                            "id": "clause",
                                            "language": "en",
                                            "text": "clause"
                                        },
                                        {
                                            "id": "utterance",
                                            "language": "en",
                                            "text": "utterance"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "perhaps I should drop Ralph a line"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.019",
                                    "subsenses": [
                                        {
                                            "id": "idcbbcb2b7-cd09-4f98-8caa-2d5d2edcc368",
                                            "synonyms": [
                                                {
                                                    "id": "correspondence",
                                                    "language": "en",
                                                    "text": "correspondence"
                                                },
                                                {
                                                    "id": "word",
                                                    "language": "en",
                                                    "text": "word"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id4018bd48-0c3f-41c9-88eb-bad03c1d9205",
                                            "registers": [
                                                "informal"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "memo",
                                                    "language": "en",
                                                    "text": "memo"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "note",
                                            "language": "en",
                                            "text": "note"
                                        },
                                        {
                                            "id": "letter",
                                            "language": "en",
                                            "text": "letter"
                                        },
                                        {
                                            "id": "card",
                                            "language": "en",
                                            "text": "card"
                                        },
                                        {
                                            "id": "postcard",
                                            "language": "en",
                                            "text": "postcard"
                                        },
                                        {
                                            "id": "message",
                                            "language": "en",
                                            "text": "message"
                                        },
                                        {
                                            "id": "bulletin",
                                            "language": "en",
                                            "text": "bulletin"
                                        },
                                        {
                                            "id": "communication",
                                            "language": "en",
                                            "text": "communication"
                                        },
                                        {
                                            "id": "epistle",
                                            "language": "en",
                                            "text": "epistle"
                                        },
                                        {
                                            "id": "missive",
                                            "language": "en",
                                            "text": "missive"
                                        },
                                        {
                                            "id": "memorandum",
                                            "language": "en",
                                            "text": "memorandum"
                                        },
                                        {
                                            "id": "dispatch",
                                            "language": "en",
                                            "text": "dispatch"
                                        },
                                        {
                                            "id": "report",
                                            "language": "en",
                                            "text": "report"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Noun",
                    "text": "line"
                },
                {
                    "entries": [
                        {
                            "homographNumber": "101",
                            "senses": [
                                {
                                    "examples": [
                                        {
                                            "text": "her face was lined with age"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.029",
                                    "synonyms": [
                                        {
                                            "id": "furrow",
                                            "language": "en",
                                            "text": "furrow"
                                        },
                                        {
                                            "id": "wrinkle",
                                            "language": "en",
                                            "text": "wrinkle"
                                        },
                                        {
                                            "id": "crease",
                                            "language": "en",
                                            "text": "crease"
                                        },
                                        {
                                            "id": "mark_with_lines",
                                            "language": "en",
                                            "text": "mark with lines"
                                        },
                                        {
                                            "id": "cover_with_lines",
                                            "language": "en",
                                            "text": "cover with lines"
                                        },
                                        {
                                            "id": "crinkle",
                                            "language": "en",
                                            "text": "crinkle"
                                        },
                                        {
                                            "id": "pucker",
                                            "language": "en",
                                            "text": "pucker"
                                        },
                                        {
                                            "id": "corrugate",
                                            "language": "en",
                                            "text": "corrugate"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "the driveway was lined by poplars"
                                        }
                                    ],
                                    "id": "t_en_gb0008794.030",
                                    "synonyms": [
                                        {
                                            "id": "border",
                                            "language": "en",
                                            "text": "border"
                                        },
                                        {
                                            "id": "edge",
                                            "language": "en",
                                            "text": "edge"
                                        },
                                        {
                                            "id": "fringe",
                                            "language": "en",
                                            "text": "fringe"
                                        },
                                        {
                                            "id": "bound",
                                            "language": "en",
                                            "text": "bound"
                                        },
                                        {
                                            "id": "skirt",
                                            "language": "en",
                                            "text": "skirt"
                                        },
                                        {
                                            "id": "hem",
                                            "language": "en",
                                            "text": "hem"
                                        },
                                        {
                                            "id": "rim",
                                            "language": "en",
                                            "text": "rim"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "homographNumber": "300",
                            "senses": [
                                {
                                    "examples": [
                                        {
                                            "text": "a cardboard box lined with a blanket"
                                        }
                                    ],
                                    "id": "t_en_gb0008795.001",
                                    "subsenses": [
                                        {
                                            "id": "ida631ef97-5225-4956-b29f-798687e32f78",
                                            "synonyms": [
                                                {
                                                    "id": "paper",
                                                    "language": "en",
                                                    "text": "paper"
                                                },
                                                {
                                                    "id": "decorate",
                                                    "language": "en",
                                                    "text": "decorate"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id38ed37a6-8565-41ce-b340-24943ff00383",
                                            "synonyms": [
                                                {
                                                    "id": "stuff",
                                                    "language": "en",
                                                    "text": "stuff"
                                                },
                                                {
                                                    "id": "fill",
                                                    "language": "en",
                                                    "text": "fill"
                                                },
                                                {
                                                    "id": "pack",
                                                    "language": "en",
                                                    "text": "pack"
                                                },
                                                {
                                                    "id": "pad",
                                                    "language": "en",
                                                    "text": "pad"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "idf79924f4-41a1-4de2-9271-9a0bd3e49758",
                                            "registers": [
                                                "archaic"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "ceil",
                                                    "language": "en",
                                                    "text": "ceil"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "cover",
                                            "language": "en",
                                            "text": "cover"
                                        },
                                        {
                                            "id": "put_a_lining_in",
                                            "language": "en",
                                            "text": "put a lining in"
                                        },
                                        {
                                            "id": "back",
                                            "language": "en",
                                            "text": "back"
                                        },
                                        {
                                            "id": "put_a_backing_on",
                                            "language": "en",
                                            "text": "put a backing on"
                                        },
                                        {
                                            "id": "interline",
                                            "language": "en",
                                            "text": "interline"
                                        },
                                        {
                                            "id": "face",
                                            "language": "en",
                                            "text": "face"
                                        },
                                        {
                                            "id": "panel",
                                            "language": "en",
                                            "text": "panel"
                                        },
                                        {
                                            "id": "inlay",
                                            "language": "en",
                                            "text": "inlay"
                                        },
                                        {
                                            "id": "reinforce",
                                            "language": "en",
                                            "text": "reinforce"
                                        },
                                        {
                                            "id": "encase",
                                            "language": "en",
                                            "text": "encase"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Verb",
                    "text": "line"
                }
            ],
            "type": "headword",
            "word": "line"
        }
    ]
}`)
		expected := []string{"underline", "underscore", "stroke", "slash", "virgule"}
		var response OxfordResponse
		json.Unmarshal(body, &response)
		actual := response.GetSynonyms()
		assertStringArray(t, expected, actual)
	})
	
	t.Run("Less than 5 in first subSense", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "metadata": {
        "provider": "Oxford University Press"
    },
    "results": [
        {
            "id": "book",
            "language": "en",
            "lexicalEntries": [
                {
                    "entries": [
                        {
                            "homographNumber": "000",
                            "senses": [
                                {
                                    "examples": [
                                        {
                                            "text": "he published his first book in 1610"
                                        }
                                    ],
                                    "id": "t_en_gb0001604.001",
                                    "subsenses": [
                                        {
                                            "id": "id39cbd3ab-ccf9-437f-8033-9d76f3565e8d",
                                            "synonyms": [
                                                {
                                                    "id": "novel",
                                                    "language": "en",
                                                    "text": "novel"
                                                },
                                                {
                                                    "id": "storybook",
                                                    "language": "en",
                                                    "text": "storybook"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id615851a2-0cdd-4616-ae5f-8859a5990257",
                                            "synonyms": [
                                                {
                                                    "id": "manual",
                                                    "language": "en",
                                                    "text": "manual"
                                                },
                                                {
                                                    "id": "handbook",
                                                    "language": "en",
                                                    "text": "handbook"
                                                },
                                                {
                                                    "id": "guide",
                                                    "language": "en",
                                                    "text": "guide"
                                                },
                                                {
                                                    "id": "companion",
                                                    "language": "en",
                                                    "text": "companion"
                                                },
                                                {
                                                    "id": "reference_book",
                                                    "language": "en",
                                                    "text": "reference book"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "idb98f6ac8-f38d-4c99-8fd0-5ac602843abf",
                                            "synonyms": [
                                                {
                                                    "id": "paperback",
                                                    "language": "en",
                                                    "text": "paperback"
                                                },
                                                {
                                                    "id": "hardback",
                                                    "language": "en",
                                                    "text": "hardback"
                                                },
                                                {
                                                    "id": "softback",
                                                    "language": "en",
                                                    "text": "softback"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "idac3bfa66-789c-474f-8a1c-f51d5017febc",
                                            "registers": [
                                                "historical"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "yellowback",
                                                    "language": "en",
                                                    "text": "yellowback"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "volume",
                                            "language": "en",
                                            "text": "volume"
                                        },
                                        {
                                            "id": "tome",
                                            "language": "en",
                                            "text": "tome"
                                        },
                                        {
                                            "id": "work",
                                            "language": "en",
                                            "text": "work"
                                        },
                                        {
                                            "id": "printed_work",
                                            "language": "en",
                                            "text": "printed work"
                                        },
                                        {
                                            "id": "publication",
                                            "language": "en",
                                            "text": "publication"
                                        },
                                        {
                                            "id": "title",
                                            "language": "en",
                                            "text": "title"
                                        },
                                        {
                                            "id": "opus",
                                            "language": "en",
                                            "text": "opus"
                                        },
                                        {
                                            "id": "treatise",
                                            "language": "en",
                                            "text": "treatise"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "he scribbled a few notes in his book"
                                        }
                                    ],
                                    "id": "t_en_gb0001604.002",
                                    "subsenses": [
                                        {
                                            "id": "id2dbcea5e-5b3f-466f-8488-35e1a8899adc",
                                            "synonyms": [
                                                {
                                                    "id": "ledger",
                                                    "language": "en",
                                                    "text": "ledger"
                                                },
                                                {
                                                    "id": "record_book",
                                                    "language": "en",
                                                    "text": "record book"
                                                },
                                                {
                                                    "id": "log",
                                                    "language": "en",
                                                    "text": "log"
                                                },
                                                {
                                                    "id": "logbook",
                                                    "language": "en",
                                                    "text": "logbook"
                                                },
                                                {
                                                    "id": "chronicle",
                                                    "language": "en",
                                                    "text": "chronicle"
                                                },
                                                {
                                                    "id": "journal",
                                                    "language": "en",
                                                    "text": "journal"
                                                },
                                                {
                                                    "id": "diary",
                                                    "language": "en",
                                                    "text": "diary"
                                                },
                                                {
                                                    "id": "daybook",
                                                    "language": "en",
                                                    "text": "daybook"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "idf6d17ce8-1b87-4fe6-8386-ad9e0ce763c4",
                                            "regions": [
                                                "British"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "jotter",
                                                    "language": "en",
                                                    "text": "jotter"
                                                },
                                                {
                                                    "id": "pocketbook",
                                                    "language": "en",
                                                    "text": "pocketbook"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id0f00d888-133d-4a35-8cc5-7337078051d7",
                                            "regions": [
                                                "North American"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "scratch_pad",
                                                    "language": "en",
                                                    "text": "scratch pad"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id6a0b06eb-81e4-4663-b74c-25383636ad18",
                                            "synonyms": [
                                                {
                                                    "id": "cahier",
                                                    "language": "en",
                                                    "text": "cahier"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "notepad",
                                            "language": "en",
                                            "text": "notepad"
                                        },
                                        {
                                            "id": "notebook",
                                            "language": "en",
                                            "text": "notebook"
                                        },
                                        {
                                            "id": "pad",
                                            "language": "en",
                                            "text": "pad"
                                        },
                                        {
                                            "id": "memo_pad",
                                            "language": "en",
                                            "text": "memo pad"
                                        },
                                        {
                                            "id": "exercise_book",
                                            "language": "en",
                                            "text": "exercise book"
                                        },
                                        {
                                            "id": "binder",
                                            "language": "en",
                                            "text": "binder"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "the council had to balance its books"
                                        }
                                    ],
                                    "id": "t_en_gb0001604.003",
                                    "subsenses": [
                                        {
                                            "id": "id779ff932-dcd7-4741-b3ca-89365ea81bfe",
                                            "synonyms": [
                                                {
                                                    "id": "account_book",
                                                    "language": "en",
                                                    "text": "account book"
                                                },
                                                {
                                                    "id": "record_book",
                                                    "language": "en",
                                                    "text": "record book"
                                                },
                                                {
                                                    "id": "ledger",
                                                    "language": "en",
                                                    "text": "ledger"
                                                },
                                                {
                                                    "id": "log",
                                                    "language": "en",
                                                    "text": "log"
                                                },
                                                {
                                                    "id": "balance_sheet",
                                                    "language": "en",
                                                    "text": "balance sheet"
                                                },
                                                {
                                                    "id": "financial_statement",
                                                    "language": "en",
                                                    "text": "financial statement"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "accounts",
                                            "language": "en",
                                            "text": "accounts"
                                        },
                                        {
                                            "id": "records",
                                            "language": "en",
                                            "text": "records"
                                        },
                                        {
                                            "id": "archives",
                                            "language": "en",
                                            "text": "archives"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Noun",
                    "text": "book"
                },
                {
                    "entries": [
                        {
                            "homographNumber": "001",
                            "senses": [
                                {
                                    "examples": [
                                        {
                                            "text": "Steven booked a table at their favourite restaurant"
                                        }
                                    ],
                                    "id": "t_en_gb0001604.005",
                                    "subsenses": [
                                        {
                                            "id": "ida0cfe413-ea3f-45d9-9542-3049fa7846a9",
                                            "synonyms": [
                                                {
                                                    "id": "charter",
                                                    "language": "en",
                                                    "text": "charter"
                                                },
                                                {
                                                    "id": "hire",
                                                    "language": "en",
                                                    "text": "hire"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "idf604502c-c10b-4858-8d52-30519f119ca6",
                                            "registers": [
                                                "informal"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "bag",
                                                    "language": "en",
                                                    "text": "bag"
                                                }
                                            ]
                                        },
                                        {
                                            "id": "id9f041001-49ce-4fc1-b17d-c629ddd7a9fc",
                                            "registers": [
                                                "dated"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "engage",
                                                    "language": "en",
                                                    "text": "engage"
                                                },
                                                {
                                                    "id": "bespeak",
                                                    "language": "en",
                                                    "text": "bespeak"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "reserve",
                                            "language": "en",
                                            "text": "reserve"
                                        },
                                        {
                                            "id": "make_a_reservation_for",
                                            "language": "en",
                                            "text": "make a reservation for"
                                        },
                                        {
                                            "id": "arrange_in_advance",
                                            "language": "en",
                                            "text": "arrange in advance"
                                        },
                                        {
                                            "id": "prearrange",
                                            "language": "en",
                                            "text": "prearrange"
                                        },
                                        {
                                            "id": "arrange_for",
                                            "language": "en",
                                            "text": "arrange for"
                                        },
                                        {
                                            "id": "order",
                                            "language": "en",
                                            "text": "order"
                                        }
                                    ]
                                },
                                {
                                    "examples": [
                                        {
                                            "text": "we booked a number of events in the Wellington Festival"
                                        }
                                    ],
                                    "id": "t_en_gb0001604.006",
                                    "subsenses": [
                                        {
                                            "id": "id333b53f9-80bf-49ba-962e-6d50e9adbe09",
                                            "regions": [
                                                "North American"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "slate",
                                                    "language": "en",
                                                    "text": "slate"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "arrange",
                                            "language": "en",
                                            "text": "arrange"
                                        },
                                        {
                                            "id": "programme",
                                            "language": "en",
                                            "text": "programme"
                                        },
                                        {
                                            "id": "schedule",
                                            "language": "en",
                                            "text": "schedule"
                                        },
                                        {
                                            "id": "timetable",
                                            "language": "en",
                                            "text": "timetable"
                                        },
                                        {
                                            "id": "line_up",
                                            "language": "en",
                                            "text": "line up"
                                        },
                                        {
                                            "id": "secure",
                                            "language": "en",
                                            "text": "secure"
                                        },
                                        {
                                            "id": "fix_up",
                                            "language": "en",
                                            "text": "fix up"
                                        },
                                        {
                                            "id": "lay_on",
                                            "language": "en",
                                            "text": "lay on"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Verb",
                    "text": "book"
                }
            ],
            "type": "headword",
            "word": "book"
        }
    ]
}`)
		expected := []string{"novel", "storybook", "manual", "handbook", "guide"}
		var response OxfordResponse
		json.Unmarshal(body, &response)
		actual := response.GetSynonyms()
		assertStringArray(t, expected, actual)
	})

	t.Run("1 synonym in subSense", func(t *testing.T) {
		var body []byte
		body = []byte(`{
    "metadata": {
        "provider": "Oxford University Press"
    },
    "results": [
        {
            "id": "refugee",
            "language": "en",
            "lexicalEntries": [
                {
                    "entries": [
                        {
                            "homographNumber": "000",
                            "senses": [
                                {
                                    "examples": [
                                        {
                                            "text": "she had fled to England as a refugee"
                                        }
                                    ],
                                    "id": "t_en_gb0012268.001",
                                    "subsenses": [
                                        {
                                            "id": "id32aa7571-48ed-4bd6-aeb6-35c53db96736",
                                            "regions": [
                                                "Australian"
                                            ],
                                            "registers": [
                                                "informal"
                                            ],
                                            "synonyms": [
                                                {
                                                    "id": "reffo",
                                                    "language": "en",
                                                    "text": "reffo"
                                                }
                                            ]
                                        }
                                    ],
                                    "synonyms": [
                                        {
                                            "id": "displaced_person",
                                            "language": "en",
                                            "text": "displaced person"
                                        },
                                        {
                                            "id": "dp",
                                            "language": "en",
                                            "text": "DP"
                                        },
                                        {
                                            "id": "escapee",
                                            "language": "en",
                                            "text": "escapee"
                                        },
                                        {
                                            "id": "fugitive",
                                            "language": "en",
                                            "text": "fugitive"
                                        },
                                        {
                                            "id": "asylum_seeker",
                                            "language": "en",
                                            "text": "asylum seeker"
                                        },
                                        {
                                            "id": "runaway",
                                            "language": "en",
                                            "text": "runaway"
                                        },
                                        {
                                            "id": "exile",
                                            "language": "en",
                                            "text": "exile"
                                        },
                                        {
                                            "id": "%C3%A9migr%C3%A9",
                                            "language": "en",
                                            "text": "émigré"
                                        },
                                        {
                                            "id": "stateless_person",
                                            "language": "en",
                                            "text": "stateless person"
                                        },
                                        {
                                            "id": "outcast",
                                            "language": "en",
                                            "text": "outcast"
                                        },
                                        {
                                            "id": "returnee",
                                            "language": "en",
                                            "text": "returnee"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "language": "en",
                    "lexicalCategory": "Noun",
                    "text": "refugee"
                }
            ],
            "type": "headword",
            "word": "refugee"
        }
    ]
}`)
		expected := []string{"reffo"}
		var response OxfordResponse
		json.Unmarshal(body, &response)
		actual := response.GetSynonyms()
		assertStringArray(t, expected, actual)
	})
}

func assertStringArray(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Fatalf("Expected %d but got %d", len(expected), len(actual))
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %+v but got %+v", expected, actual)
	}
}

