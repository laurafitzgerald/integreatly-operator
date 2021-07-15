package addon

import "strings"

var quotaConfig = `
[{
	"name": "50M",
	"rate-limiting": {
		"unit": "minute",
		"requests_per_unit": 34723,
		"alert_limits": []
	},
	"resources": {
		"backend_listener": {
			"replicas": 5,
			"resources": {
				"requests": {
					"cpu": 0.5,
					"memory": "470Mi"
				},
				"limits": {
					"cpu": 1,
					"memory": "520Mi"
				}
			}
		},
		"backend_worker": {
			"replicas": 4,
			"resources": {
				"requests": {
					"cpu": 0.4,
					"memory": "100Mi"
				},
				"limits": {
					"cpu": 0.6,
					"memory": "100Mi"
				}
			}
		},
		"apicast_production": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.6,
					"memory": "275Mi"
				},
				"limits": {
					"cpu": 1,
					"memory": "300Mi"
				}
			}
		},
	"rhssouser": {
		"replicas": 3,
		"resources": {
			"requests": {
				"cpu": 1,
				"memory": "2000Mi"
			},
			"limits": {
				"cpu": 2,
				"memory": "2000Mi"
			}
		}
	},
	"ratelimit": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.15,
					"memory": "50Mi"
				},
				"limits": {
					"cpu": 0.3,
					"memory": "100Mi"
				}
			}
		}
	}
},
{
	"name": "20M",
	"rate-limiting": {
		"unit": "minute",
		"requests_per_unit": 13889,
		"alert_limits": []
	},
	"resources": {
		"backend_listener": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.25,
					"memory": "450Mi"
				},
				"limits": {
					"cpu": 0.6,
					"memory": "500Mi"
				}
			}
		},
		"backend_worker": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.15,
					"memory": "100Mi"
				},
				"limits": {
					"cpu": 0.3,
					"memory": "100Mi"
				}
			}
		},
		"apicast_production": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.3,
					"memory": "250Mi"
				},
				"limits": {
					"cpu": 0.6,
					"memory": "300Mi"
				}
			}
		},
	"rhssouser": {
		"replicas": 3,
		"resources": {
			"requests": {
				"cpu": 0.75,
				"memory": "1500Mi"
			},
			"limits": {
				"cpu": 1.5,
				"memory": "1500Mi"
			}
		}
	},
	"ratelimit": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.10,
					"memory": "50Mi"
				},
				"limits": {
					"cpu": 0.20,
					"memory": "100Mi"
				}
			}
		}
	}
},
	 {
	"name": "10M",
	"rate-limiting": {
		"unit": "minute",
		"requests_per_unit": 6945,
		"alert_limits": []
	},
	"resources": {
		"backend_listener": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.15,
					"memory": "450Mi"
				},
				"limits": {
					"cpu": 0.5,
					"memory": "500Mi"
				}
			}
		},
		"backend_worker": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.1,
					"memory": "100Mi"
				},
				"limits": {
					"cpu": 0.25,
					"memory": "100Mi"
				}
			}
		},
		"apicast_production": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.2,
					"memory": "250Mi"
				},
				"limits": {
					"cpu": 0.5,
					"memory": "300Mi"
				}
			}
		},
	"rhssouser": {
		"replicas": 3,
		"resources": {
			"requests": {
				"cpu": 0.75,
				"memory": "1500Mi"
			},
			"limits": {
				"cpu": 1.5,
				"memory": "1500Mi"
			}
		}
	},
	"ratelimit": {
			"replicas": 3,
			"resources": {
				"requests": {
					"cpu": 0.05,
					"memory": "50Mi"
				},
				"limits": {
					"cpu": 0.15,
					"memory": "100Mi"
				}
			}
		}
	}
},
{
  "name": "5M",
  "rate-limiting": {
	  "unit": "minute",
	  "requests_per_unit": 3473,
	  "alert_limits": []
  },
  "resources": {
	  "backend_listener": {
		  "replicas": 3,
		  "resources": {
			  "requests": {
				  "cpu": 0.1,
				  "memory": "450Mi"
			  },
			  "limits": {
				  "cpu": 0.3,
				  "memory": "500Mi"
			  }
		  }
	  },
	  "backend_worker": {
		  "replicas": 3,
		  "resources": {
			  "requests": {
				  "cpu": 0.05,
				  "memory": "100Mi"
			  },
			  "limits": {
				  "cpu": 0.15,
				  "memory": "100Mi"
			  }
		  }
	  },
	  "apicast_production": {
		  "replicas": 3,
		  "resources": {
			  "requests": {
				  "cpu": 0.1,
				  "memory": "250Mi"
			  },
			  "limits": {
				  "cpu": 0.3,
				  "memory": "300Mi"
			  }
		  }
	  },
	  "rhssouser": {
		  "replicas": 3,
		  "resources": {
			  "requests": {
				  "cpu": 0.75,
				  "memory": "1500Mi"
			  },
			  "limits": {
				  "cpu": 1.5,
				  "memory": "1500Mi"
			  }
		  }
	  },
	  "ratelimit": {
		  "replicas": 3,
		  "resources": {
			   "requests": {
				   "cpu": 0.05,
				   "memory": "50Mi"
			   },
			   "limits": {
				   "cpu": 0.15,
				   "memory": "100Mi"
			   }
		   }
	  }
  }
},
{
	"name": "1M",
	"rate-limiting": {
		"unit": "minute",
		"requests_per_unit": 695,
		"alert_limits": []
	},
	"resources": {
		"backend_listener": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.06,
					"memory": "450Mi"
				},
				"limits": {
					"cpu": 0.18,
					"memory": "500Mi"
				}
			}
		},
		"backend_worker": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.03,
					"memory": "60Mi"
				},
				"limits": {
					"cpu": 0.09,
					"memory": "100Mi"
				}
			}
		},
		"apicast_production": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.06,
					"memory": "250Mi"
				},
				"limits": {
					"cpu": 0.18,
					"memory": "300Mi"
				}
			}
		},
	  "rhssouser": {
		   "replicas": 2,
		   "resources": {
			   "requests": {
				   "cpu": 0.75,
				   "memory": "1500Mi"
			   },
			   "limits": {
				   "cpu": 1.5,
				   "memory": "1500Mi"
			   }
		   }
	   },
		"ratelimit": {
			"replicas": 2,
            "resources": {
			  "requests": {
				  "cpu": 0.02,
				  "memory": "40Mi"
			  },
			  "limits": {
				  "cpu": 0.06,
				  "memory": "80Mi"
			  }
		  }
		}
	}
},
{
	"name": "0E",
	"rate-limiting": {
		"unit": "minute",
		"requests_per_unit": 70,
		"alert_limits": []
	},
	"resources": {
		"backend_listener": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.06,
					"memory": "450Mi"
				},
				"limits": {
					"cpu": 0.18,
					"memory": "500Mi"
				}
			}
		},
		"backend_worker": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.03,
					"memory": "60Mi"
				},
				"limits": {
					"cpu": 0.09,
					"memory": "100Mi"
				}
			}
		},
		"apicast_production": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.06,
					"memory": "250Mi"
				},
				"limits": {
					"cpu": 0.18,
					"memory": "300Mi"
				}
			}
		},
	  "rhssouser": {
		   "replicas": 2,
		   "resources": {
			   "requests": {
				   "cpu": 0.75,
				   "memory": "1500Mi"
			   },
			   "limits": {
				   "cpu": 1.5,
				   "memory": "1500Mi"
			   }
		   }
	   },
		"ratelimit": {
			"replicas": 2,
            "resources": {
			  "requests": {
				  "cpu": 0.02,
				  "memory": "40Mi"
			  },
			  "limits": {
				  "cpu": 0.06,
				  "memory": "80Mi"
			  }
		  }
		}
	}
},
{
	"name": "100K",
	"rate-limiting": {
		"unit": "minute",
		"requests_per_unit": 70,
		"alert_limits": []
	},
	"resources": {
		"backend_listener": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.06,
					"memory": "450Mi"
				},
				"limits": {
					"cpu": 0.18,
					"memory": "500Mi"
				}
			}
		},
		"backend_worker": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.03,
					"memory": "60Mi"
				},
				"limits": {
					"cpu": 0.09,
					"memory": "100Mi"
				}
			}
		},
		"apicast_production": {
			"replicas": 2,
			"resources": {
				"requests": {
					"cpu": 0.06,
					"memory": "250Mi"
				},
				"limits": {
					"cpu": 0.18,
					"memory": "300Mi"
				}
			}
		},
	  "rhssouser": {
		   "replicas": 2,
		   "resources": {
			   "requests": {
				   "cpu": 0.75,
				   "memory": "1500Mi"
			   },
			   "limits": {
				   "cpu": 1.5,
				   "memory": "1500Mi"
			   }
		   }
	   },
		"ratelimit": {
			"replicas": 2,
            "resources": {
			  "requests": {
				  "cpu": 0.02,
				  "memory": "40Mi"
			  },
			  "limits": {
				  "cpu": 0.06,
				  "memory": "80Mi"
			  }
		  }
		}
	}
}]
`

func GetQuotaConfig() string {
	return strings.TrimSpace(quotaConfig)
}
