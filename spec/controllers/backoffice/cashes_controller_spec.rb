# frozen_string_literal: true

RSpec.describe Backoffice::CashesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { is_expected.to eq :collection }
    end

    context do
      before do
        allow(subject).to receive(:params).and_return(:params)

        allow(Cash).to receive(:order).with(:name).and_return(:scope)

        allow(CashSearcher).to receive(:search).with(:scope, :params).and_return(:collection)
      end

      its(:collection) { is_expected.to eq :collection }
    end
  end

  describe '#resource' do
    context 'when @resource is set' do
      before { subject.instance_variable_set(:@resource, :resource) }

      its(:resource) { is_expected.to eq :resource }
    end

    context 'when @resource is not set' do
      let(:params) { { id: 52 } }

      before do
        allow(subject).to receive(:params).and_return(params)

        allow(Cash).to receive(:find).with(52).and_return(:resource)
      end

      its(:resource) { is_expected.to eq :resource }
    end
  end

  describe '#initialize_resource' do
    before do
      allow(Cash).to receive(:new).and_return(:resource)

      subject.send(:initialize_resource)
    end

    its(:resource) { is_expected.to eq :resource }
  end

  describe '#resource_params' do
    before { allow(subject).to receive(:params).and_return(params) }

    %w[uah usd eur].each do |currency|
      context "when currency is `#{ currency }`" do
        let(:params) { acp(currency:, cash: { name: nil, formula: nil, supercategory: nil, favorite: nil }) }

        its(:resource_params) { is_expected.to eq params.require(:cash).permit!.merge(currency:) }
      end
    end
  end

  describe '#build_resource' do
    before do
      allow(subject).to receive(:resource_params).and_return(:resource_params)

      allow(Cash).to receive(:new).with(:resource_params).and_return(:resource)

      subject.send(:build_resource)
    end

    its(:resource) { is_expected.to eq :resource }
  end
end
