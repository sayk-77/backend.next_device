import http from 'k6/http';
import { check, group } from 'k6';
import { sleep } from 'k6';

export let options = {
    stages: [
        { duration: '5s', target: 1000 },
        { duration: '10s', target: 5000 },
        { duration: '5s', target: 0 },
    ],
};

export default function () {
    group('Category Test', function () {
        const categories = ['laptop', 'mobile', 'tablet'];
        const category = categories[Math.floor(Math.random() * categories.length)];

        const url = `http://localhost:5000/api/catalog/${category}`;


        const res = http.get(url);

        check(res, {
            'status is 200': (r) => r.status === 200,
            'response time is less than 1s': (r) => r.timings.duration < 1000,
        });
    });

    group('Product Filter Test', function () {
        const filters = {
            brands: ['Apple'],
            cameraQualities: [],
            memories: [],
            os: [],
            priceFrom: 0,
            priceTo: 1000000,
            ram: ['8'],
            screenFrom: 0,
            screenTo: 15,
        };

        const url = 'http://localhost:5000/api/product/mobile/query';


        const res = http.post(url, JSON.stringify(filters), {
            headers: { 'Content-Type': 'application/json' },
        });


        check(res, {
            'status is 200': (r) => r.status === 200,
            'response contains products': (r) => r.json().length > 0,
        });
    });

    sleep(1);
}
