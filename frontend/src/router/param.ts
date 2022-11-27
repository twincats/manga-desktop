import { RouteLocationNormalized } from 'vue-router'

/**
 * Casts props into proper data types.
 * Props ending in 'ID' and the prop 'id' are cast to Number automatically.
 * To cast other props or override the defaults, pass a mapping like this:
 * @example
 * {
 *  path: '/:isNice/:age/:hatSize',
 *  name: 'foo route',
 *  props: typedProps({ isNice: Boolean, age: Number, hatSize: Number}),
 * },
 * @param {Object} mapping
 * @returns
 */
export const typedProps = (mapping: Record<string, any>) => {
  return (route: Record<string, any>) => {
    let nameType = Object.entries(mapping) // [[param1, Number], [param2, String]]
    let nameRouteParam = nameType.map(([name, fn]) => [
      name,
      fn(route.params[name]),
    ]) // [[param1, 1], [param2, "hello"]]
    let props = Object.fromEntries(nameRouteParam) // {param1: 1, param2: "hello"}
    return props
  }
}
