require 'rails_helper'

RSpec.describe ItemsController, type: :controller do
  describe '#index.js' do
    before { get :index, xhr: true, format: :js }

    it { should render_template :index }
  end

  describe '#create.js' do
    let(:item) { double }

    let(:params) do
      { item: { date: '2014-04-22', formula: '2+2', category_id: '1', description: 'Buys' } }
    end

    before { expect(Item).to receive(:new).with(permit! params[:item]).and_return(item) }

    context do
      before { expect(item).to receive(:save).and_return(true) }

      before { post :create, params: params, format: :js }

      it { should render_template :create }
    end

    context do
      before { expect(item).to receive(:save).and_return(false) }

      before { post :create, params: params, format: :js }

      it { should render_template :new }
    end
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

  describe '#resource' do
    before { expect(subject).to receive(:params).and_return({ id: 43 }) }

    before { expect(Item).to receive(:find).with(43).and_return(:resource) }

    its(:resource) { should eq :resource }
  end

  describe '#update.js' do
    let(:item) { double }

    let(:params) do
      { item: { date: '2014-04-22', formula: '2+2', category_id: '1', description: 'Buys' }, id: 1 }
    end

    before { subject.instance_variable_set :@resource, item }

    before { expect(item).to receive(:update!).with(permit! params[:item]) }

    before { put :update, params: params, format: :js }

    it { should render_template :update }
  end

  describe '#destroy.js' do
    let(:item) { double }

    before { subject.instance_variable_set :@resource, item }

    before { expect(item).to receive(:destroy) }

    before { delete :destroy, params: { id: 13 }, format: :js }

    it { should render_template :destroy }
  end
end
