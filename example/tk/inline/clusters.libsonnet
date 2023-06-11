[
  {
	// import the microservices example
    local tempo = import '../tempo-microservices/main.jsonnet',

    name: 'test-drp-kcc-to-tanzu',
    apiServer: 'https://0.0.0.0:6443',
    namespace: 'tempo-grafana',
    
    data: tempo,

    dataOverride: {
      _images+:: {
        // images can be overridden here if desired
      },

      _config+:: {

        // config can be overridden here if desired
        
      },

    },

  },
]
