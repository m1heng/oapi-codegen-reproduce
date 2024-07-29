# oapi-codegen-reproduce
missing `prompt` field in response

in openapi.yml
paths:

```
  /api/public/v2/prompts/{promptName}:
    get:
      description: Get a prompt
      operationId: prompts_get
      tags:
        - Prompts
      parameters:
        - name: promptName
          in: path
          description: The name of the prompt
          required: true
          schema:
            type: string
        - name: version
          in: query
          description: Version of the prompt to be retrieved.
          required: false
          schema:
            type: integer
            nullable: true
        - name: label
          in: query
          description: >-
            Label of the prompt to be retrieved. Defaults to "production" if no
            label or version is set.
          required: false
          schema:
            type: string
            nullable: true
      responses:
        '200':
          description: ''
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Prompt'
        '400':
          description: ''
          content:
            application/json:
              schema: {}
        '401':
          description: ''
          content:
            application/json:
              schema: {}
        '403':
          description: ''
          content:
            application/json:
              schema: {}
        '404':
          description: ''
          content:
            application/json:
              schema: {}
        '405':
          description: ''
          content:
            application/json:
              schema: {}
      security: *ref_0
```

// components
```
    Prompt:
      title: Prompt
      oneOf:
        - type: object
          allOf:
            - type: object
              properties:
                type:
                  type: string
                  enum:
                    - chat
            - $ref: '#/components/schemas/ChatPrompt'
          required:
            - type
        - type: object
          allOf:
            - type: object
              properties:
                type:
                  type: string
                  enum:
                    - text
            - $ref: '#/components/schemas/TextPrompt'
          required:
            - type
    BasePrompt:
      title: BasePrompt
      type: object
      properties:
        name:
          type: string
        version:
          type: integer
        config: {}
        labels:
          type: array
          items:
            type: string
          description: List of deployment labels of this prompt version.
        tags:
          type: array
          items:
            type: string
          description: >-
            List of tags. Used to filter via UI and API. The same across
            versions of a prompt.
      required:
        - name
        - version
        - config
        - labels
        - tags
    ChatMessage:
      title: ChatMessage
      type: object
      properties:
        role:
          type: string
        content:
          type: string
      required:
        - role
        - content
    TextPrompt:
      title: TextPrompt
      type: object
      properties:
        prompt:
          type: string
      required:
        - prompt
      allOf:
        - $ref: '#/components/schemas/BasePrompt'
    ChatPrompt:
      title: ChatPrompt
      type: object
      properties:
        prompt:
          type: array
          items:
            $ref: '#/components/schemas/ChatMessage'
      required:
        - prompt
      allOf:
        - $ref: '#/components/schemas/BasePrompt'
```

generated code
```
// Prompt0 defines model for .
type Prompt0 struct {
	Config interface{} `json:"config"`

	// Labels List of deployment labels of this prompt version.
	Labels []string `json:"labels"`
	Name   string   `json:"name"`

	// Tags List of tags. Used to filter via UI and API. The same across versions of a prompt.
	Tags    []string     `json:"tags"`
	Type    *Prompt0Type `json:"type,omitempty"`
	Version int          `json:"version"`
}

// Prompt0Type defines model for Prompt.0.Type.
type Prompt0Type string

// Prompt1 defines model for .
type Prompt1 struct {
	Config interface{} `json:"config"`

	// Labels List of deployment labels of this prompt version.
	Labels []string `json:"labels"`
	Name   string   `json:"name"`

	// Tags List of tags. Used to filter via UI and API. The same across versions of a prompt.
	Tags    []string     `json:"tags"`
	Type    *Prompt1Type `json:"type,omitempty"`
	Version int          `json:"version"`
}
```