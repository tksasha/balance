require 'rails_helper'

RSpec.describe ItemsController, type: :controller do
  describe '#index.js' do
    before { get :index, xhr: true, format: :js }

    it { should render_template :index }
  end

  describe '#items' do
    let(:item) { stub_model Item }

    let(:date) { Date.today }

    let(:date_range) { date.beginning_of_month..date.end_of_month }

    before do
      #
      # Item.search(date_range, nil).includes(:category)
      #
      expect(Item).to receive(:search).with(date_range, nil) do
        double.tap do |a|
          expect(a).to receive(:includes).with(:category)
        end
      end
    end

    it { expect { subject.send :items, date_range }.to_not raise_error }
  end
end
