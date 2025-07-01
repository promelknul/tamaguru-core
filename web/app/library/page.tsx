'use client'
import dynamic from 'next/dynamic'
import { useEffect, useState } from 'react'
const FG = dynamic(() => import('react-force-graph-3d'), { ssr: false })
export default function Library() {
  const [data, setData] = useState({ nodes: [], links: [] })
  useEffect(() => { fetch('/graph.json').then(r => r.json()).then(setData) }, [])
  return <FG graphData={data} onNodeClick={n => window.location = `/library?node=${n.id}`} />
}
