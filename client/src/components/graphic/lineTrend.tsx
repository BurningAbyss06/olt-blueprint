import SpinnerBasicComponent from '../spinner/basic';
import { Chart } from 'chart.js/auto';
import { useEffect } from 'react';
import { LoadStatus } from '../../constant/loadStatus';
import { getUnit } from '../../utils/transform';
import { getMonthString } from '../../utils/date';
import type { MeasurementSchema } from '../../schemas/measurement';
import type { TrendGraphSchema } from '../../schemas/trend';
import type { LoadingStateValue } from '../../types/loadingState';
import React from 'react';

/**
 * @interface Data required for the line graph.
 * 
 * @param {LoadingStateValue} loading Loading state of the graph component.
 * @param {string} title Title of the graph.
 * @param {string} canvasID ID of the canvas.
 * @param {TrendGraphSchema[]} data Data of the graph.
 */
interface LineProps {
  loading: LoadingStateValue;
  title: string;
  canvasID: string;
  data: TrendGraphSchema[];
}

export default function LineTrendGraphComponent(content: LineProps) {

  /**
   * Message when no search has been done.
   */
  const ViewWaitingData = () => {
    return(<h1 className="text-2xl font-semibold text-gray-300 text-center">Sin búsqueda</h1>);
  }
  
  /**
   * Message when there is no data to show.
   */
  const ViewWithoutData = () => {
    return(<h1 className="text-2xl font-semibold text-gray-300 text-center">No se encontró información para gráficar</h1>);
  }

  /**
   * Render the chart when data is available.
   */
  const ViewData = () => {
    return (
      <>
        <h1 className="self-start text-2xl font-semibold text-gray-700">{content.title}</h1>
        <canvas id={content.canvasID}></canvas>
      </>
    )
  }

  useEffect(() => {
    const canvas = document.getElementById(content.canvasID) as HTMLCanvasElement;
    if (canvas) {
      new Chart(canvas, {
        type: 'line',
        data: {
          labels: content.data.map((trend: TrendGraphSchema) => {
            return `${getMonthString(trend.month)}`;
          }),
          datasets: [
            {
              label: 'In',
              data: content.data.map((trend: TrendGraphSchema) => trend.in),
              fill: true,
              borderColor: 'rgb(205, 7, 7)',
              backgroundColor: 'rgba(205, 7, 7, 0.4)',
              tension: 0.1
            },
            {
              label: 'Out',
              data: content.data.map((trend: TrendGraphSchema) => trend.out),
              fill: true,
              borderColor: 'rgb(32, 35, 229)',
              backgroundColor: 'rgba(32, 35, 229, 0.4)',
              tension: 0.1
            }
          ]
        },
        options: {
          plugins: {
              legend: {
                  position: 'bottom',
                  align: 'center',
              }
          },
          scales: {
            x: { display: true },
            y: {
              ticks: {
                callback: (value) => {
                  return getUnit(Number(value));
                },
              },
            },
          },
        }
      });
    }
  }, [content.data]);

  return (
    <div className="w-full lg:w-5/6 max-w-full min-h-96 px-4 py-6 bg-white rounded-xl flex flex-col justify-center items-center gap-4">
      {content.loading === LoadStatus.EMPTY && ViewWaitingData()}
      {content.loading === LoadStatus.LOADING && <SpinnerBasicComponent />}
      {content.loading === LoadStatus.LOADED && content.data && content.data.length > 0 && ViewData()}
      {content.loading === LoadStatus.LOADED && content.data && content.data.length <= 0 && ViewWithoutData()}
      {content.loading === LoadStatus.LOADED && !content.data && ViewWithoutData()}
    </div>
  );
}
