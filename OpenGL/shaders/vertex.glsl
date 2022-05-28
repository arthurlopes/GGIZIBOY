#version 460

in vec3 vp;
in float cp;
out float c;

void main() {
    gl_Position = vec4(vp, 1.0);
    c = cp;
}