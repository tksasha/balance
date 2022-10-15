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
end
