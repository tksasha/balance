# frozen_string_literal: true

RSpec.describe CashesController do
  describe '#resource_params' do
    let(:params) { acp(cash: { name: nil, formula: nil }) }

    before { allow(subject).to receive(:params).and_return(params) }

    its(:resource_params) { is_expected.to eq params.require(:cash).permit! }
  end

  describe '#resource' do
    context 'when @resource is set' do
      before { subject.instance_variable_set(:@resource, :resource) }

      its(:resource) { is_expected.to eq :resource }
    end

    context 'when @resource is not set' do
      let(:params) { { id: 21 } }

      before do
        allow(subject).to receive(:params).and_return(params)

        allow(Cash).to receive(:find).with(21).and_return(:resource)
      end

      its(:resource) { is_expected.to eq :resource }
    end
  end
end
