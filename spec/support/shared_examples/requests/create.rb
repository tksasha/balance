# frozen_string_literal: true

RSpec.shared_examples 'create.js' do
  it { is_expected.to render_template :create }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }

  it { expect(response).to have_http_status(:created) }
end
