import * as v from 'valibot'

export const validationSchema = v.object({
  database: v.optional(v.string()),
  host: v.pipe(v.string(), v.trim(), v.nonEmpty('* required')),
  port: v.pipe(
    v.string(),
    v.trim(),
    v.nonEmpty('* required'),
    v.transform(Number),
    v.integer('* port be an integer'),
    v.minValue(1, '* port must be ≥ 1'),
    v.maxValue(65535, '* port must be ≤ 65535') // uint16 related to RFC
  ),
  user: v.pipe(v.string(), v.trim(), v.nonEmpty('* required')),
  password: v.pipe(v.string(), v.trim(), v.nonEmpty('* required')),
  sslmode: v.optional(v.string(), '')
})

export type FormSubmitData = v.InferOutput<typeof validationSchema>
