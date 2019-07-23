<script>
  import { onMount } from 'svelte'
  import { max, min } from 'd3-array'
  import { scaleLinear, scaleTime } from 'd3-scale'

  export let users = []
  const pointsPossible = 400
  let wonders = []
  let userPoints = {}

  onMount(async function fetchUserWonders() {
    const res = await fetch('/api/v1/user-wonders')
    const body = await res.json()
    wonders = body.data
    users.forEach(user => {
      userPoints[user.name] = convertWondersToPoints(wonders.filter(w => w.userId === user.id), pointsPossible)
    })
  })

  function convertWondersToPoints(wonders, pointsPossible) {
    return wonders.map(w => ({
      x: new Date(w.created),
      y: --pointsPossible
    }))
  }

  function formatDisplayDate(str) {
    const [yyyy, mm, dd] = str.split('T')[0].split('-').map(n => parseInt(n, 10))
    const date = new Date(yyyy, mm - 1, dd)
    const month = date.toLocaleString('en-us', { month: 'short' })
    return dd + ' ' + month
  }

  function flatten(userPoints) {
    return Object.values(userPoints).reduce((acc, points) => [...acc, ...points], [])
  }

  function color(i) {
    const colors = [
      '#FFBF00',
      '#EF8532',
      '#D1342B',
      '#549D3E'
    ]
    return colors[i % colors.length]
  }

  const padding = { top: 20, right: 15, bottom: 20, left: 25 };

  let width = 500;
  let height = 200;

  // TODO: a line for each user
  $: points = userPoints.dallin || []

  $: minX = min(flatten(userPoints), p => p.x)
  $: maxX = max(flatten(userPoints), p => p.x)
  $: xScale = scaleTime()
          .domain([minX, maxX])
          .range([padding.left, width - padding.right]);

  $: minY = min(flatten(userPoints), p => p.y)
  $: maxY = max(flatten(userPoints), p => p.y)
  $: yScale = scaleLinear()
          .domain([minY, maxY])
          .range([height - padding.bottom, padding.top]);

  $: xTicks = [minX, maxX]
  $: yTicks = [minY, pointsPossible]

  $: paths = Object.values(userPoints).map(points => `M${points.map(p => `${xScale(p.x)},${yScale(p.y)}`).join('L')}`)

  function formatDate (date) {
    if (!date) return
    const month = date.toLocaleString('en-us', { month: 'short' })
    return date.getDate() + ' ' + month
  }
</script>

<style>
svg {
		position: relative;
		width: 100%;
		height: 200px;
		overflow: visible;
	}

	.tick {
		font-size: .725em;
		font-weight: 200;
	}

	.tick line {
		stroke: #efefef;
		stroke-dasharray: 2;
	}

	.tick text {
		fill: #fff;
		text-anchor: start;
	}

	.tick.tick-300 line {
		stroke-dasharray: 0;
	}

	.x-axis .tick text {
		text-anchor: middle;
	}

	.path-line {
		fill: none;
                stroke: #FFBF00;
		stroke-linejoin: round;
		stroke-linecap: round;
		stroke-width: 2;
	}

        .users {
          display: flex;
          justify-content: space-between;
        }
        .users li {
          display: flex;
          align-items: center;
        }
        .color {
          height: 1rem;
          width: 2rem;
          margin-right: 0.25rem;
        }

</style>

<div class="chart" bind:clientWidth={width} bind:clientHeight={height}>
	<svg>
		<!-- y axis -->
		<g class="axis y-axis" transform="translate(0, {padding.top})">
			{#each yTicks as tick}
				<g class="tick tick-{tick}" transform="translate(0, {yScale(tick) - padding.bottom})">
					<line x2="100%"></line>
					<text y="-4">{tick} {tick === pointsPossible ? ' points' : ''}</text>
				</g>
			{/each}
		</g>

		<!-- x axis -->
		<g class="axis x-axis">
			{#each xTicks as tick}
				<g class="tick tick-{ tick }" transform="translate({xScale(tick)},{height})">
					<line y1="-{height}" y2="-{padding.bottom}" x1="0" x2="0"></line>
					<text y="-2">{formatDate(tick)}</text>
				</g>
			{/each}
		</g>

		<!-- data -->
                {#each paths as path, i}
                  <path class="path-line" style="stroke:{color(i)}" d={path}></path>
                {/each}
	</svg>
</div>

<ul class="users">
  {#each users as user, i}
    <li>
      <div class="color" style="background:{color(i)}"></div>
      {user.name}
    </li>
  {/each}
</ul>