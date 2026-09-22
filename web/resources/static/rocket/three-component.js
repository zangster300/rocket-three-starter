import { rocket } from 'datastar-rocket';
import * as THREE from 'three';
import { OrbitControls } from 'three/addons/controls/OrbitControls.js';


rocket('three-component', {
    mode: 'light',
    props: ({ number }) => ({
        height: number.clamp(1, 100).default(0),
        width: number.clamp(1, 100).default(0),
        depth: number.clamp(1, 100).default(0),
        x: number.clamp(-10, 10).default(0),
        y: number.clamp(-10, 10).default(0),
        z: number.clamp(-10, 10).default(0),
    }),
    onFirstRender: ({ refs, props }) => {
        console.log(props)

        let ro

        function main() {
            // init
            const canvas = refs.canvas;
            const renderer = new THREE.WebGLRenderer({ antialias: true, canvas });

            // camera
            const fov = 75;
            const aspect = 2; // the canvas default
            const near = 0.1;
            const far = 10;
            const camera = new THREE.PerspectiveCamera(fov, aspect, near, far);
            camera.position.z = 6;

            const controls = new OrbitControls(camera, renderer.domElement);
            controls.update();

            // scene
            const scene = new THREE.Scene();

            {
                // light
                const color = 0xffffff;
                const intensity = 3;
                const light = new THREE.DirectionalLight(color, intensity);
                light.position.set(-1, 2, 4);
                scene.add(light);
            }

            // box geometry
            const boxWidth = props.width;
            const boxHeight = props.height;
            const boxDepth = props.depth;
            const geometry = new THREE.BoxGeometry(boxWidth, boxHeight, boxDepth);

            // cube instance creator function
            function makeInstance(geometry, color, x) {
                const material = new THREE.MeshPhongMaterial({ color });

                const cube = new THREE.Mesh(geometry, material);
                scene.add(cube);

                cube.position.x = x;

                return cube;
            }

            const cubes = [
                makeInstance(geometry, 0x8844aa, -2),
                makeInstance(geometry, 0x44aa88, 0),
                makeInstance(geometry, 0xaa8844, 2),
            ];

            // responsive
            // handles HD-DPI displays
            function resizeRendererToDisplaySize(renderer) {
                const canvas = renderer.domElement;
                const pixelRatio = Math.min(window.devicePixelRatio, 2);
                const width = Math.floor(window.innerWidth * pixelRatio);
                const height = Math.floor(window.innerHeight * pixelRatio);
                const needResize = canvas.width !== width || canvas.height !== height;
                if (needResize) {
                    renderer.setSize(width, height, false);
                }

                return needResize;
            }

            ro = new ResizeObserver(() => {
                resizeRendererToDisplaySize(renderer)
            })
            ro.observe(canvas.parentElement || canvas)

            // render loop
            function render(time) {
                time *= 0.001;

                controls.update();

                if (resizeRendererToDisplaySize(renderer)) {
                    const canvas = renderer.domElement;
                    camera.aspect = canvas.clientWidth / canvas.clientHeight;
                    camera.updateProjectionMatrix();
                }

                cubes.forEach((cube, ndx) => {
                    const speed = 1 + ndx * 0.1;
                    const rot = time * speed;
                    cube.rotation.x = rot;
                    cube.rotation.y = rot;
                    if (ndx === 1) {
                        cube.position.y = props.y
                    } else {
                        cube.position.y = -props.y
                    }
                });
                

                renderer.render(scene, camera);

                requestAnimationFrame(render);
            }

            requestAnimationFrame(render);
        }

        main();
    },
    render: ({ html }) => html`
    <style>
        .three-container {
            color: var(--color-base-content);
            height: 100%;
            width: 100%;
        }

        canvas {
            height: 100%;
            width: 100%;
            display: block;
        }

        .debug {
            position: absolute;
            color: var(--color-base-content);
            padding: 1rem;
        }
    </style>
    <div class="three-container">
        <canvas data-ref:canvas data-ignore-morph></canvas>
    </div>
  `,
})