/**
 * A function for tagged template literals that encodes params for URLS.
 */
export function encodeParams(template, ...expressions) {
    return template.reduce((accumulator, part, i) => {
        return accumulator + encodeURIComponent(expressions[i - 1]) + part;
    });
}
