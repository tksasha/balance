# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#current_month' do
    let(:month) { Month.new(2021, 3) }

    before { allow(subject).to receive(:month).and_return(month) }

    its(:current_month) { is_expected.to eq 3 }
  end
end
