#version 460

out vec4 frag_colour;
in float c;

void main() {
    frag_colour = vec4(c, c, c, 1.0);
}