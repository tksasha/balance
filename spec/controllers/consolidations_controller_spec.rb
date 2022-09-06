# frozen_string_literal: true

RSpec.describe ConsolidationsController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { is_expected.to eq :collection }
    end

    context do
      before do
        allow(subject).to receive(:params).and_return(:params)

        allow(Consolidations::GetCollectionService).to receive(:call).with(:params).and_return(:collection)
      end

      its(:collection) { is_expected.to eq :collection }
    end
  end
end
