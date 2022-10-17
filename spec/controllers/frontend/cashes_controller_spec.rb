# frozen_string_literal: true

RSpec.describe Frontend::CashesController do
  describe '#collection' do
    context 'when @collection is set' do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { is_expected.to eq :collection }
    end

    context 'when @collection is not set' do
      before do
        allow(subject).to receive(:params).and_return(:params)

        allow(Cash).to receive(:order).with(:name).and_return(:cashes)

        allow(CashSearcher).to receive(:search).with(:cashes, :params).and_return(:collection)
      end

      its(:collection) { is_expected.to eq :collection }
    end
  end

  describe '#resource' do
    context 'when @resource is set' do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { is_expected.to eq :resource }
    end

    context 'when @resource is not set' do
      let(:params) { { id: 57 } }

      before do
        allow(subject).to receive(:params).and_return(params)

        allow(Cash).to receive(:find).with(57).and_return(:resource)
      end

      its(:resource) { is_expected.to eq :resource }
    end
  end

  describe '#resource_params' do
    let(:params) { acp(cash: { name: nil, formula: nil }) }

    before { allow(subject).to receive(:params).and_return(params) }

    its(:resource_params) { is_expected.to eq params.require(:cash).permit! }
  end
end
