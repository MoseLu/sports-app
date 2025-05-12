export default {
  framework: {
    config: {
      notify: {
        position: 'top',
        timeout: 3000,
        textColor: 'white',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
        classes: 'q-mt-xl',
        style: {
          'max-width': '400px',
          margin: '0 auto',
          'border-radius': '8px',
          'box-shadow': '0 4px 12px rgba(0, 0, 0, 0.1)',
        },
      },
    },
  },
};
