# frozen_string_literal: true

RSpec.describe Backoffice::CashesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { is_expected.to eq :collection }
    end

    context do
      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Cashes::GetCollectionService).to receive(:call).with(:params).and_return(:collection) }

      its(:collection) { is_expected.to eq :collection }
    end
  end

  describe '#result' do
    context do
      before { subject.instance_variable_set :@result, :result }

      its(:result) { is_expected.to eq :result }
    end

    context do
      before { allow(subject).to receive(:action_name).and_return(:action_name) }

      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Cashes::GetResultService).to receive(:call).with(:action_name, :params).and_return(:result) }

      its(:result) { is_expected.to eq :result }
    end
  end
end
