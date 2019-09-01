# frozen_string_literal: true

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

  describe '#resource_params' do
    let :params do
      acp item: { date: nil, formula: nil, category_id: nil, description: nil }
    end

    before { expect(subject).to receive(:params).and_return(params) }

    its(:resource_params) { should eq params[:item].permit! }
  end

  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      before { expect(subject).to receive(:params).and_return(id: 26) }

      before { expect(Item).to receive(:find).with(26).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:resource_params) }

    before { expect(Item).to receive(:new).with(:resource_params).and_return(:resource) }

    before { subject.send :build_resource }

    its(:resource) { should eq :resource }
  end

  it_behaves_like :create, format: :js do
    let(:success) { -> { should render_template(:create).with_status(201) } }

    let(:failure) { -> { should render_template(:new).with_status(422) } }
  end

  it_behaves_like :update, format: :js do
    let(:success) { -> { should render_template(:update).with_status(200) } }

    let(:failure) { -> { should render_template(:edit).with_status(422) } }
  end

  it_behaves_like :destroy, format: :js do
    let(:success) { -> { should render_template(:destroy).with_status(200) } }
  end
end
